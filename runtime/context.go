package runtime

import (
	"errors"
	"time"
)

const NoTimeout time.Duration = 0
const DefaultTimeout = 5 * time.Second

type Context struct {
	parentContext *Context
	variables     map[string]Value
	functions     [16]*Function

	result  Value
	err     error
	timeout *time.Timer
	UI      UIInterface
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

func NewReturn(res Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		c.result, c.err = res.Call(c)
		return c.result, c.err
	})
}

func NewGetVariable(vn string) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		val, ok := c.variables[vn]
		if ok {
			return val, nil
		}
		return Int(0), nil
	})
}

func NewSetVariable(vn string, value Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		val, err := value.Call(c)
		if err != nil {
			return nil, err
		}
		c.variables[vn] = val
		return val, nil
	})
}