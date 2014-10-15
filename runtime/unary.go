package runtime

import (
	"../token"
)

func NewNegInt(value *Callable, p *token.Token) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		iv, err := callInt(value, c)
		if err != nil {
			return nil, err
		}

		return -iv, nil
	}, value)
}

func NewNegBool(value *Callable, p *token.Token) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		bv, err := callBool(value, c)
		if err != nil {
			return nil, err
		}

		return !bv, nil
	}, value)
}
