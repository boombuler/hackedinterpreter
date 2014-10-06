package termwnd

import (
	"github.com/nsf/termbox-go"
	"sync"
)

type Handler interface {
	Input(ev termbox.Event)
}

var (
	windows       []*Window
	activeWindow  *Window
	qryUpdChan    chan *Window
	quitChan      chan struct{}
	setActWndChan chan *Window
	wgexit        *sync.WaitGroup
)

func NewWindow(handler Handler) *Window {
	w := new(Window)
	w.handler = handler
	w.clear(termbox.Size())
	if windows == nil {
		windows = make([]*Window, 0)
		activeWindow = w
	}
	windows = append(windows, w)
	return w
}

func Start() {
	wgexit = new(sync.WaitGroup)
	wgexit.Add(1)
	termbox.Init()
	eventChan := make(chan termbox.Event)

	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
		close(eventChan)
	}()

	qryUpdChan = make(chan *Window)
	setActWndChan = make(chan *Window)
	quitChan = make(chan struct{})

	go func() {
		defer wgexit.Done()
		defer termbox.Close()
		for {
			select {
			case <-quitChan:
				return
			case ev := <-eventChan:
				if ev.Type == termbox.EventResize {
					w, h := termbox.Size()
					for _, wnd := range windows {
						wnd.clear(w, h)
					}
				} else if ev.Type == termbox.EventKey {
					go activeWindow.handler.Input(ev)
				} else {
					return
				}
			case w := <-setActWndChan:
				activeWindow = w
				w.render()
			case wq := <-qryUpdChan:
				if wq == activeWindow {
					wq.render()
				}
			}
		}

	}()
}

func Wait() {
	wgexit.Wait()
}

func Stop() {
	quitChan <- struct{}{}
}
