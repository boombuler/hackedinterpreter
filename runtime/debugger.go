package runtime

import (
	"errors"
)

type executor func(c *Callable, ctx *Context) (Value, error)

// A simple executor function which can be replaced by a Debugger Fn
func simpleExecutor(c *Callable, ctx *Context) (Value, error) {
	return c.fn(ctx)
}

type Debugger struct {
	curCtx      *Context
	mainCtx     *Context
	curCallable *Callable
}

func AttachDebugger(c *Context) (*Debugger, error) {
	if c == nil || c.parentContext != nil {
		return nil, errors.New("Debugger can only be attached to main context")
	}
	d := &Debugger{
		mainCtx: c,
	}
	c.executor = d.exec
	return d, nil
}

func (d *Debugger) exec(c *Callable, ctx *Context) (Value, error) {
	d.curCtx = ctx
	d.curCallable = c
	// Check here for BreakPoint etc

	return c.fn(ctx)
}
