package runtime

import (
	"../token"
	"errors"
	"time"
)

const NoTimeout time.Duration = 0
const DefaultTimeout = 5 * time.Second

type Context struct {
	parentContext *Context
	variables     map[string]Value
	functions     [16]*Function

	result   Value
	err      error
	timeout  *time.Timer
	UI       UIInterface
	executor executor
}

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

func NewContext(d time.Duration) *Context {
	c := new(Context)
	c.variables = make(map[string]Value)
	c.executor = simpleExecutor
	if d > 0 {
		c.timeout = time.NewTimer(d)
	}
	return c
}

func (c *Context) newChildContext() *Context {
	return &Context{
		variables:     make(map[string]Value),
		parentContext: c,
	}
}

func (c *Context) SetInput(value Value) {
	c.variables["input"] = value
}

func (ctx *Context) Call(c Callable) (Value, error) {
	if ctx.err != nil || ctx.result != nil {
		return ctx.result, ctx.err
	}
	cc := ctx
	for cc.parentContext != nil {
		cc = cc.parentContext
	}

	val, err := cc.executor(c, ctx)

	if err != nil {
		re, ok := err.(*RuntimeError)
		meta := c.Meta()
		if ok {
			if meta.Token != nil {
				re.CallStack = append(re.CallStack, &meta.Pos)
			}
			return nil, re
		} else {
			return nil, &RuntimeError{
				Message:   err.Error(),
				Position:  &meta.Pos,
				CallStack: make([]*token.Pos, 0),
			}
		}
	}
	return val, nil
}

func (c *Context) ui() UIInterface {
	for c.parentContext != nil {
		if c.UI != nil {
			return c.UI
		}
		c = c.parentContext
	}
	return c.UI
}

func (c *Context) forceExit() bool {
	if c.err != nil || c.result != nil {
		return true
	}
	if c.parentContext != nil {
		return c.parentContext.forceExit()
	} else if c.timeout != nil {
		select {
		case <-c.timeout.C:
			c.err = errors.New("timeout")
			c.result = nil
			return true
		default:
			return false
		}
	} else {
		return false
	}
}

func NewReturn(res Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		c.result, c.err = c.Call(res)
		return c.result, c.err
	}, res)
}
