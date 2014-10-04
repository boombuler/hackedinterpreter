package runtime

import (
	"../token"
	"errors"
)

type DebuggerExecutionMode byte

const (
	Continue DebuggerExecutionMode = iota
	Step
)

type BreakEvent struct {
	Continue chan<- DebuggerExecutionMode
	Token    *token.Token
}

type executor func(c *Callable, ctx *Context) (Value, error)

// A simple executor function which can be replaced by a Debugger Fn
func simpleExecutor(c *Callable, ctx *Context) (Value, error) {
	return c.fn(ctx)
}

type Debugger struct {
	curCtx      *Context
	mainCtx     *Context
	curCallable *Callable

	mode DebuggerExecutionMode

	breakEventChan chan<- *BreakEvent
	codeBPs        map[*token.Token]bool // change bool to something else if conditional BPs should be impl.
}

func AttachDebugger(c *Context, be chan<- *BreakEvent) (*Debugger, error) {
	if c == nil || c.parentContext != nil {
		return nil, errors.New("Debugger can only be attached to main context")
	}
	d := &Debugger{
		mainCtx:        c,
		breakEventChan: be,
		mode:           Step,
	}
	c.executor = d.exec
	return d, nil
}

func (d *Debugger) ToggleCodeBreakPoint(t *token.Token) {
	_, ok := d.codeBPs[t]
	if ok {
		delete(d.codeBPs, t)
	} else {
		d.codeBPs[t] = true
	}

}

func (d *Debugger) exec(c *Callable, ctx *Context) (Value, error) {
	d.curCtx = ctx
	d.curCallable = c
	// Check here for BreakPoint etc
	if d.mode == Step || d.codeBPs[c.Token] {
		// We reached a BP...
		if d.breakEventChan != nil {
			res := make(chan DebuggerExecutionMode)
			d.breakEventChan <- &BreakEvent{res, c.Token}
			d.mode = <-res
		}
	}
	return c.fn(ctx)
}
