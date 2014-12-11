package runtime

import (
	"github.com/boombuler/hackedinterpreter/token"
)

func NewNegInt(value Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		iv, err := callInt(value, c)
		if err != nil {
			return nil, err
		}

		return -iv, nil
	}, value)
}

func NewNegBool(value Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		bv, err := callBool(value, c)
		if err != nil {
			return nil, err
		}

		return !bv, nil
	}, value)
}

func NewIncAssignable(a Assignable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		iv, err := callInt(a, c)
		if err != nil {
			return iv, err
		}
		return iv, a.Set(iv+1, c)
	}, a)
}

func NewDecAssignable(a Assignable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		iv, err := callInt(a, c)
		if err != nil {
			return iv, err
		}
		return iv, a.Set(iv-1, c)
	}, a)
}
