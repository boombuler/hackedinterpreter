package main

import (
	"./lexer"
	"./runtime"
	"./token"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

type dbgWorkspace struct {
	lexer *lexer.DebugLexer
	dbg   *runtime.Debugger
	code  *runtime.Callable

	commandEB    *EditBox
	curBreakEv   *runtime.BreakEvent
	columnOffset int
	lineOffset   int
	maxLine      int
	maxCol       int
}

func (ws *dbgWorkspace) renderCode(x, y, w, h int) {
	for _, t := range ws.lexer.Tokens {
		if t.Line <= ws.lineOffset || t.Line > h+ws.lineOffset {
			continue
		}

		bgColor := termbox.ColorBlack
		fgColor := getForeground(t)
		if ws.dbg.IsCodeBreakPoint(t) {
			bgColor = termbox.ColorRed
		}

		if ws.curBreakEv != nil && ws.curBreakEv.Token == t {
			fgColor, bgColor = bgColor, fgColor
		}

		text := string(t.Lit)
		c := t.Column - len(text) - ws.columnOffset
		l := t.Line - ws.lineOffset - 1
		if c < 0 {
			if len(text) <= -c {
				continue
			} else {
				text = text[-c:]
				c = 0
			}
		}

		textOut(text, c+x, l+y, w-c, fgColor, bgColor)
	}
}

func textOut(s string, x, y, w int, fg, bg termbox.Attribute) {
	w += x
	for _, r := range s {
		if x < w {
			termbox.SetCell(x, y, r, fg, bg)
			x += 1
		}
	}
}

func (ws *dbgWorkspace) renderVars(x, y, w, h int) {
	for xx := x; xx < x+w; xx++ {
		for yy := y; yy < y+h; yy++ {
			termbox.SetCell(xx, yy, ' ', termbox.ColorBlack, termbox.ColorWhite)
		}
	}
	x += 1
	w -= 1

	vars := ws.dbg.GetVars()
	varNames := []string{"input"}
	for x := rune(97); x < rune(112); x++ {
		varNames = append(varNames, "var_"+string(x))
	}

	for _, varN := range varNames {
		val, ok := vars[varN]
		if !ok {
			val = 0
		}
		valStr := runtime.ToString(val)

		textOut(varN+":", x, y, w, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
		textOut(valStr, x+7, y, w-7, termbox.ColorBlack, termbox.ColorWhite)
		y += 1
	}
}

func (ws *dbgWorkspace) renderLineNos(x, y, h int) int {
	lineNoLen := len(fmt.Sprintf("%d", ws.maxLine))
	lineNo := ws.lineOffset + 1
	format := fmt.Sprintf("%%%dd", lineNoLen)
	for ; y < h; y++ {
		textOut(fmt.Sprintf(format, lineNo), x, y, lineNoLen, termbox.ColorYellow, termbox.ColorBlack)
		lineNo += 1
		if lineNo > ws.maxLine {
			break
		}
	}
	return lineNoLen
}

func invalidateOffset(curVal, maxVal, space int) int {
	if curVal < 0 {
		return 0
	}
	maxOffset := maxVal - space
	if curVal > maxOffset {
		if maxOffset < 0 {
			return 0
		}
		return maxOffset
	}
	return curVal
}

func (ws *dbgWorkspace) render(scrollToToken *token.Token) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()
	varWndSize := w / 3

	ws.renderVars(w-varWndSize, 0, varWndSize, h)
	if ws.commandEB != nil {
		h -= 1
		ws.commandEB.Draw(0, h, w-varWndSize, 1)

	}

	if scrollToToken != nil {
		if scrollToToken.Line <= ws.lineOffset {
			ws.lineOffset = scrollToToken.Line - 1
		}
		if scrollToToken.Line > ws.lineOffset+h {
			ws.lineOffset = scrollToToken.Line - h
		}
	}

	ws.lineOffset = invalidateOffset(ws.lineOffset, ws.maxLine, h)
	lineNoWidth := ws.renderLineNos(0, 0, h) + 1
	codeWidth := w - varWndSize - lineNoWidth

	if scrollToToken != nil {
		if scrollToToken.Column <= ws.columnOffset {
			ws.columnOffset = scrollToToken.Column - 1
		}
		if scrollToToken.Column > ws.columnOffset+codeWidth {
			ws.columnOffset = scrollToToken.Column - codeWidth
		}
	}

	ws.columnOffset = invalidateOffset(ws.columnOffset, ws.maxCol, codeWidth)

	ws.renderCode(lineNoWidth, 0, codeWidth, h)
	termbox.Flush()
}

func (ws *dbgWorkspace) handleKey(ev termbox.Event) {
	if ws.commandEB != nil {
		if ws.commandEB.handleKey(ev) {
			cmd := string(ws.commandEB.text)
			ws.handleCommand(cmd)
			ws.commandEB = nil
		}
	} else {
		switch ev.Key {
		case termbox.KeyArrowDown:
			ws.lineOffset += 1
		case termbox.KeyArrowUp:
			ws.lineOffset -= 1
		case termbox.KeyArrowRight:
			ws.columnOffset += 1
		case termbox.KeyArrowLeft:
			ws.columnOffset -= 1
		case termbox.KeyEnter:
			ws.commandEB = &EditBox{
				Foreground: termbox.ColorWhite,
				Background: termbox.ColorBlue,
			}
		default:
			switch ev.Ch {
			case 's':
				if be := ws.curBreakEv; be != nil {
					ws.curBreakEv = nil
					be.Continue <- runtime.Step
				}

			case 'c':
				if be := ws.curBreakEv; be != nil {
					ws.curBreakEv = nil
					be.Continue <- runtime.Continue
				}
			}

		}
	}
}

func startDebugCode(c *runtime.Callable, input runtime.Value) (*runtime.Debugger, <-chan *runtime.BreakEvent, <-chan runtime.Value, <-chan error) {
	ctx := runtime.NewContext(runtime.NoTimeout)
	ctx.SetInput(input)
	bpEv := make(chan *runtime.BreakEvent)
	valRes := make(chan runtime.Value)
	errRes := make(chan error)

	dbg, _ := runtime.AttachDebugger(ctx, bpEv)

	go func() {
		defer close(bpEv)
		defer close(valRes)
		defer close(errRes)
		val, err := c.Call(ctx)
		if err != nil {
			errRes <- err
		}
		valRes <- val
	}()

	return dbg, bpEv, valRes, errRes
}

func RunDebugger(fileName string, input runtime.Value) {
	code, lex, err := fromFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = termbox.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	termClosed := false
	closeTerm := func() {
		if !termClosed {
			termbox.Close()
			termClosed = true
		}
	}
	defer closeTerm()

	maxLine := 0
	maxCol := 0
	for _, t := range lex.Tokens {
		if t.Line > maxLine {
			maxLine = t.Line
		}
		mc := t.Column + len(t.Lit)
		if mc > maxCol {
			maxCol = mc
		}
	}

	debugger, breakChan, valueChan, errorChan := startDebugCode(code, input)
	shutdownChan := make(chan struct{})
	eventChan := make(chan termbox.Event)
	go func() {
		defer close(shutdownChan)
		defer close(eventChan)
		for {
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				shutdownChan <- struct{}{}
				return
			} else {
				eventChan <- ev
			}
		}
	}()

	ws := &dbgWorkspace{
		lexer:        lex,
		dbg:          debugger,
		code:         code,
		columnOffset: 0,
		lineOffset:   0,
		maxLine:      maxLine,
		maxCol:       maxCol,
	}

	var stt *token.Token = nil
	for {
		ws.render(stt)
		stt = nil
		select {
		case val := <-valueChan:
			closeTerm()
			fmt.Fprintln(os.Stdout, val)
			return
		case err := <-errorChan:
			closeTerm()
			fmt.Fprintln(os.Stderr, err)
			return
		case <-shutdownChan:
			return
		case ev := <-eventChan:
			if ev.Type == termbox.EventKey {
				ws.handleKey(ev)
			}
		case be := <-breakChan:
			ws.curBreakEv = be
			stt = ws.curBreakEv.Token
		}
	}
}
