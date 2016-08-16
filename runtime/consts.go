package runtime

import (
	"github.com/boombuler/hackedinterpreter/token"
	"strconv"
)

func NewConstBool(val bool, p *token.Token) (Callable, error) {
	return newCallable(p, func(c *Context) (Value, error) {
		return val, nil
	}), nil
}

func NewConstInt(text string, p *token.Token) (Callable, error) {
	i, err := strconv.Atoi(text)
	if err != nil {
		return nil, err
	}
	return newCallable(p, func(c *Context) (Value, error) {
		return i, nil
	}), nil
}
