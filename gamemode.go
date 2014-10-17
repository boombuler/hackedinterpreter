package main

import (
	"./runtime"
	"./termwnd"
	"github.com/nsf/termbox-go"
)

type GameWindow struct {
	*termwnd.Window
	up    int
	down  int
	left  int
	right int
	a     int
	b     int
}

func (ui *GameWindow) BtnA() (bool, error) {
	return ui.a > 0, nil
}
func (ui *GameWindow) BtnB() (bool, error) {
	return ui.b > 0, nil
}
func (ui *GameWindow) Down() (bool, error) {
	return ui.down > 0, nil
}
func (ui *GameWindow) Up() (bool, error) {
	return ui.up > 0, nil
}
func (ui *GameWindow) Right() (bool, error) {
	return ui.right > 0, nil
}
func (ui *GameWindow) Left() (bool, error) {
	return ui.left > 0, nil
}
func (ui *GameWindow) Width() (int, error) {
	w := ui.Window.Width()
	return w, nil
}
func (ui *GameWindow) Height() (int, error) {
	h := ui.Window.Height()
	return h, nil
}
func (ui *GameWindow) Draw(x, y int) error {
	return ui.SetCell(x, y, ' ', termbox.ColorGreen, termbox.ColorGreen)
}
func (ui *GameWindow) DrawText(x, y int, text string) error {
	for _, r := range text {
		err := ui.SetCell(x, y, r, termbox.ColorGreen, termbox.ColorDefault)
		if err != nil {
			return err
		}
		x += 1
	}
	return nil
}

func (ui *GameWindow) Input(ev termbox.Event) {
	if ev.Ch == 'a' {
		ui.a += 1
	} else if ev.Ch == 'b' {
		ui.b += 1
	} else {
		switch ev.Key {
		case termbox.KeyArrowDown:
			ui.down += 1
		case termbox.KeyArrowUp:
			ui.up += 1
		case termbox.KeyArrowRight:
			ui.right += 1
		case termbox.KeyArrowLeft:
			ui.left += 1
		case termbox.KeyEsc:
			termwnd.Stop()
		}
	}
}

func NewGameWindow(code runtime.Callable, ctx *runtime.Context) *GameWindow {
	ui := new(GameWindow)
	ui.Window = termwnd.NewWindow(ui)
	ctx.UI = ui
	go func() {
		for {
			u, d, l, r, a, b := ui.up, ui.down, ui.left, ui.right, ui.a, ui.b
			ui.Clear()
			_, err := ctx.Call(code)
			if err != nil {
				panic(err.Error())
			}
			ui.Render()
			ui.up -= u
			ui.down -= d
			ui.left -= l
			ui.right -= r
			ui.a -= a
			ui.b -= b
		}
	}()
	return ui
}

func RunGame(code runtime.Callable) {
	c := runtime.NewContext(runtime.NoTimeout)
	NewGameWindow(code, c)
}
