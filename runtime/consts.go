package runtime

import (
	"strconv"
)

func NewConstBool(val bool) (*Callable, error) {
	return newCallable(nil, func(c *Context) (Value, error) {
		return val, nil
	}), nil
}

func NewConstInt(text string) (*Callable, error) {
	i, err := strconv.Atoi(text)
	if err != nil {
		return nil, err
	}
	return newCallable(nil, func(c *Context) (Value, error) {
		return i, nil
	}), nil
}
