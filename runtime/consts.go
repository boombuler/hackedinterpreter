package runtime

import (
	"../util"
)

type callableFunc func(c *Context) (Value, error)

func (cf callableFunc) Call(c *Context) (Value, error) {
	return cf(c)
}

type Bool bool

type Int int64

func (cb Int) Call(c *Context) (Value, error) {
	return cb, nil
}

func (cb Bool) Call(c *Context) (Value, error) {
	return cb, nil
}

func NewConstInt(text []byte) (Int, error) {
	i, err := util.IntValue(text)
	return Int(i), err
}
