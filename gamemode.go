package main

import (
	"./runtime"
	"github.com/nsf/termbox-go"
)

type UiTermBox struct {
	up    bool
	down  bool
	left  bool
	right bool
	a     bool
	b     bool
}

func (ui *UiTermBox) BtnA() (bool, error) {
	return ui.a, nil
}
func (ui *UiTermBox) BtnB() (bool, error) {
	return ui.b, nil
}
func (ui *UiTermBox) Down() (bool, error) {
	return ui.down, nil
}
func (ui *UiTermBox) Up() (bool, error) {
	return ui.up, nil
}
func (ui *UiTermBox) Right() (bool, error) {
	return ui.right, nil
}
func (ui *UiTermBox) Left() (bool, error) {
	return ui.left, nil
}
func (ui *UiTermBox) Width() (int, error) {
	w, _ := termbox.Size()
	return w, nil
}
func (ui *UiTermBox) Height() (int, error) {
	_, h := termbox.Size()
	return h, nil
}
func (ui *UiTermBox) Draw(x, y int) error {
	x, y = x-1, y-1
	termbox.SetCell(x, y, ' ', termbox.ColorGreen, termbox.ColorGreen)
	return nil
}
func (ui *UiTermBox) DrawText(x, y int, text string) error {
	x, y = x-1, y-1
	for _, r := range text {
		termbox.SetCell(x, y, r, termbox.ColorGreen, termbox.ColorDefault)
		x += 1
	}
	return nil
}

func gameLoop(code runtime.Callable, closeChan <-chan struct{}, input <-chan termbox.Event) {
	ctx := runtime.NewContext(runtime.NoTimeout)
	ui := new(UiTermBox)
	ctx.UI = ui
	go func() {
		for {
			select {
			case <-closeChan:
				return
			case ev := <-input:
				if ev.Ch == 'a' {
					ui.a = true
				} else if ev.Ch == 'b' {
					ui.b = true
				} else {
					switch ev.Key {
					case termbox.KeyArrowDown:
						ui.down = true
					case termbox.KeyArrowUp:
						ui.up = true
					case termbox.KeyArrowRight:
						ui.right = true
					case termbox.KeyArrowLeft:
						ui.left = true
					}
				}
			default:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				_, err := code.Call(ctx)
				if err != nil {
					termbox.Close()
					panic(err.Error())
				}
				termbox.Flush()
				ui.up = false
				ui.down = false
				ui.left = false
				ui.right = false
				ui.a = false
				ui.b = false
			}
		}
	}()
}

func RunGame(code runtime.Callable) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	closeChan := make(chan struct{})
	input := make(chan termbox.Event)
	gameLoop(code, closeChan, input)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				closeChan <- struct{}{}
				close(closeChan)
				close(input)
				return
			default:
				input <- ev
			}
		}
	}
}
