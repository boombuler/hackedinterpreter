package runtime

import (
	"strconv"
)

type callableFunc func(c *Context) (Value, error)

func (cf callableFunc) Call(c *Context) (Value, error) {
	if c.err != nil || c.result != nil {
		return c.result, c.err
	}
	return cf(c)
}

type Bool bool

type Int int64

func (cb Int) Call(c *Context) (Value, error) {
	if c.err != nil || c.result != nil {
		return c.result, c.err
	}
	return cb, nil
}

func (cb Bool) Call(c *Context) (Value, error) {
	if c.err != nil || c.result != nil {
		return c.result, c.err
	}
	return cb, nil
}

func NewConstInt(text string) (Int, error) {
	i, err := strconv.Atoi(text)
	return Int(i), err
}
