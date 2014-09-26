package main

import (
	"errors"
)

type UIInterface interface {
	BtnA() (bool, error)
	BtnB() (bool, error)
	Down() (bool, error)
	Up() (bool, error)
	Right() (bool, error)
	Left() (bool, error)
	Width() (int, error)
	Height() (int, error)
	Draw(x, y int) error
	DrawText(x, y int, text string) error
}

type UIStub struct{}

var NotSupported = errors.New("Not Supported")

func (ui *UIStub) BtnA() (bool, error) {
	return false, NotSupported
}
func (ui *UIStub) BtnB() (bool, error) {
	return false, NotSupported
}
func (ui *UIStub) Down() (bool, error) {
	return false, NotSupported
}
func (ui *UIStub) Up() (bool, error) {
	return false, NotSupported
}
func (ui *UIStub) Right() (bool, error) {
	return false, NotSupported
}
func (ui *UIStub) Left() (bool, error) {
	return false, NotSupported
}
func (ui *UIStub) Width() (int, error) {
	return 0, NotSupported
}
func (ui *UIStub) Height() (int, error) {
	return 0, NotSupported
}
func (ui *UIStub) Draw(x, y int) error {
	return NotSupported
}
func (ui *UIStub) DrawText(x, y int, text string) error {
	return NotSupported
}
