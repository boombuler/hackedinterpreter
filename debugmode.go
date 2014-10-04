package main

import (
	"./lexer"
	"./runtime"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

type dbgWorkspace struct {
	lexer *lexer.DebugLexer
	dbg   *runtime.Debugger

	curBreakEv   *runtime.BreakEvent
	columnOffset int
	lineOffset   int
	maxLine      int
}

func (ws *dbgWorkspace) render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	lastLine := 0
	lineTokenCnt := 0

	for _, t := range ws.lexer.Tokens {
		if lastLine != t.Line {
			lineTokenCnt = 0
			lastLine = t.Line
		}
		if t.Line <= ws.lineOffset || t.Line > h+ws.lineOffset {
			continue
		}

		bgColor := termbox.ColorBlack
		if ws.curBreakEv != nil && ws.curBreakEv.Token == t {
			bgColor = termbox.ColorRed
		}

		text := string(t.Lit)
		c := t.Column - len(text) - lineTokenCnt - ws.columnOffset
		l := t.Line - ws.lineOffset
		lineTokenCnt += 1
		for _, r := range text {
			if c <= ws.columnOffset || c > w+ws.columnOffset {
				continue
			}
			termbox.SetCell(c-1, l-1, r, termbox.ColorWhite, bgColor)
			c += 1
		}
	}
	termbox.Flush()
}

func (ws *dbgWorkspace) handleKey(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowDown:
		if ws.lineOffset < ws.maxLine-1 {
			ws.lineOffset += 1
		}
	case termbox.KeyArrowUp:
		if ws.lineOffset > 0 {
			ws.lineOffset -= 1
		}
	default:
		switch ev.Ch {
		case 's':
			if be := ws.curBreakEv; be != nil {
				ws.curBreakEv = nil
				be.Continue <- runtime.Step
			}
		}
	}

}

func startDebugCode(c *runtime.Callable) (*runtime.Debugger, <-chan *runtime.BreakEvent, <-chan runtime.Value, <-chan error) {
	ctx := runtime.NewContext(runtime.NoTimeout)
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

func RunDebugger(fileName string) {
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
	defer termbox.Close()

	maxLine := 0
	for _, t := range lex.Tokens {
		if t.Line > maxLine {
			maxLine = t.Line
		}
	}

	debugger, breakChan, valueChan, errorChan := startDebugCode(code)
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
		columnOffset: 0,
		lineOffset:   0,
		maxLine:      maxLine,
	}

	for {
		ws.render()

		select {
		case val := <-valueChan:
			fmt.Fprintln(os.Stdout, val)
			return
		case err := <-errorChan:
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
		}
	}
}
