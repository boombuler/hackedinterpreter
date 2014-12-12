package main

import (
	"fmt"
	"github.com/boombuler/hackedinterpreter/lexer"
	"github.com/boombuler/hackedinterpreter/runtime"
	"github.com/boombuler/hackedinterpreter/termwnd"
	"github.com/boombuler/hackedinterpreter/token"
	"github.com/nsf/termbox-go"
	"os"
)

type dbgWorkspace struct {
	*termwnd.Window
	lexer   *lexer.DebugLexer
	dbg     *runtime.Debugger
	code    runtime.Callable
	gameWnd *GameWindow

	commandEB    *termwnd.EditBox
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
			fgAttrs := fgColor & (termbox.AttrBold | termbox.AttrReverse | termbox.AttrUnderline)
			fgColor = fgColor ^ fgAttrs
			fgColor, bgColor = bgColor, fgColor
			fgColor = fgColor | fgAttrs
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

		ws.textOut(text, c+x, l+y, w-c, fgColor, bgColor)
	}
}

func (ws *dbgWorkspace) textOut(s string, x, y, w int, fg, bg termbox.Attribute) {
	w += x
	for _, r := range s {
		if x < w {
			ws.SetCell(x, y, r, fg, bg)
			x += 1
		}
	}
}

func (ws *dbgWorkspace) renderVars(x, y, w, h int) {
	for xx := x; xx < x+w; xx++ {
		for yy := y; yy < y+h; yy++ {
			ws.SetCell(xx, yy, ' ', termbox.ColorBlack, termbox.ColorWhite)
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
		valStr := runtime.ToString(val, true)

		bg := termbox.ColorWhite

		if ws.dbg.IsMemBreakPoint(varN) {
			bg = termbox.ColorRed
		}

		ws.textOut(varN, x, y, w, termbox.ColorBlack|termbox.AttrBold, bg)
		ws.textOut(":", x+5, y, w-5, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
		ws.textOut(valStr, x+7, y, w-7, termbox.ColorBlack, termbox.ColorWhite)
		y += 1
	}
}

func (ws *dbgWorkspace) renderLineNos(x, y, h int) int {
	lineNoLen := len(fmt.Sprintf("%d", ws.maxLine))
	lineNo := ws.lineOffset + 1
	format := fmt.Sprintf("%%%dd", lineNoLen)
	for ; y < h; y++ {
		ws.textOut(fmt.Sprintf(format, lineNo), x, y, lineNoLen, termbox.ColorYellow, termbox.ColorBlack)
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
	ws.Clear()
	w, h := ws.Width(), ws.Height()
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
	ws.Render()
}

func (ws *dbgWorkspace) Input(ev termbox.Event) {
	if ev.Key == termbox.KeyEsc {
		termwnd.Stop()
	}

	if ws.commandEB != nil {
		if ws.commandEB.HandleKey(ev) {
			cmd := string(ws.commandEB.Text)
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
			ws.commandEB = &termwnd.EditBox{
				Window:     ws.Window,
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
					if ws.gameWnd != nil {
						ws.gameWnd.Activate()
					}
					be.Continue <- runtime.Continue
				}
			}
		}
	}
	ws.render(nil)
}

func startDebugCode(c runtime.Callable, isGame bool, input runtime.Value) (*runtime.Debugger, <-chan *runtime.BreakEvent, <-chan runtime.Value, <-chan error, *GameWindow) {
	ctx := runtime.NewContext(runtime.NoTimeout)
	if input != nil {
		ctx.SetVariable("input", input)
	}
	bpEv := make(chan *runtime.BreakEvent)

	dbg, _ := runtime.AttachDebugger(ctx, bpEv)
	if isGame {
		gw := NewGameWindow(c, ctx)
		return dbg, bpEv, nil, nil, gw
	} else {
		valRes := make(chan runtime.Value)
		errRes := make(chan error)
		go func() {
			defer close(bpEv)
			defer close(valRes)
			defer close(errRes)
			val, err := ctx.Call(c)
			if err != nil {
				errRes <- err
			}
			valRes <- val
		}()

		return dbg, bpEv, valRes, errRes, nil
	}
}

func RunDebugger(code runtime.Callable, lex *lexer.DebugLexer, isGame bool, input runtime.Value) {
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

	ws := &dbgWorkspace{
		lexer:        lex,
		code:         code,
		columnOffset: 0,
		lineOffset:   0,
		maxLine:      maxLine,
		maxCol:       maxCol,
	}
	ws.Window = termwnd.NewWindow(ws)

	debugger, breakChan, valueChan, errorChan, gameWnd := startDebugCode(code, isGame, input)
	ws.dbg = debugger
	ws.gameWnd = gameWnd

	go func() {
		var stt *token.Token = nil
		for {
			ws.render(stt)
			stt = nil
			select {
			case val := <-valueChan:
				termwnd.Stop()
				fmt.Fprintln(os.Stdout, val)
				return
			case err := <-errorChan:
				termwnd.Stop()
				fmt.Fprintln(os.Stderr, err)
				return
			case be := <-breakChan:
				ws.curBreakEv = be
				stt = ws.curBreakEv.Token
				ws.Activate()
			}
		}
	}()
}
