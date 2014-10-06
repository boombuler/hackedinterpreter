package termwnd

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Window struct {
	buffer  [][]*termbox.Cell
	width   int
	height  int
	handler Handler
}

func (w *Window) Width() int {
	return w.width
}
func (w *Window) Height() int {
	return w.height
}

func (wnd *Window) clear(w, h int) {
	wnd.buffer = make([][]*termbox.Cell, h)
	for y := 0; y < h; y++ {
		wnd.buffer[y] = make([]*termbox.Cell, w)
		for x := 0; x < w; x++ {
			wnd.buffer[y][x] = new(termbox.Cell)
		}
	}
	wnd.width = w
	wnd.height = h
	wnd.Clear()
}

func (w *Window) Clear() {
	for _, cols := range w.buffer {
		for _, cell := range cols {
			cell.Bg = termbox.ColorDefault
			cell.Fg = termbox.ColorDefault
			cell.Ch = ' '
		}
	}
}

func (w *Window) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) error {
	if y < 0 || y >= w.height || x < 0 || x >= w.width {
		return fmt.Errorf("out of screen: x: %d / y: %d", x, y)
	}
	cell := w.buffer[y][x]
	cell.Ch = ch
	cell.Fg = fg
	cell.Bg = bg
	return nil
}

func (w *Window) Activate() {
	setActWndChan <- w
}

func (w *Window) Render() {
	qryUpdChan <- w
}

func (w *Window) render() {
	for y, cols := range w.buffer {
		for x, cell := range cols {
			termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
		}
	}
	termbox.Flush()
}
