package main

import (
	"./lexer"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

type dbgWorkspace struct {
	lexer        *lexer.DebugLexer
	columnOffset int
	lineOffset   int
}

func (ws *dbgWorkspace) drawCode() {
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

		text := string(t.Lit)
		c := t.Column - len(text) - lineTokenCnt - ws.columnOffset
		l := t.Line - ws.lineOffset
		lineTokenCnt += 1
		for _, r := range text {
			if c <= ws.columnOffset || c > w+ws.columnOffset {
				continue
			}
			termbox.SetCell(c-1, l-1, r, termbox.ColorWhite, termbox.ColorBlack)
			c += 1
		}
	}
	termbox.Flush()
}

func RunDebugger(fileName string) {
	code, lex, err := fromFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	_ = code
	err = termbox.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer termbox.Close()
	ws := &dbgWorkspace{
		lexer:        lex,
		columnOffset: 0,
		lineOffset:   0,
	}

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		ws.drawCode()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				ws.lineOffset += 1
			case termbox.KeyArrowUp:
				ws.lineOffset -= 1

			case termbox.KeyEsc:
				return
			default:
			}
		}
	}
}
