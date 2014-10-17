package runtime

import (
	"../token"
	"errors"
)

func NewMul(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return nil, err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return nil, err
		}

		return i1 * i2, nil
	}, c1, c2)
}

func NewDiv(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return nil, err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return nil, err
		}
		if i2 == 0 {
			return nil, errors.New("Division by zero!")
		}

		return i1 / i2, nil
	}, c1, c2)
}

func NewAdd(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return nil, err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return nil, err
		}

		return i1 + i2, nil
	}, c1, c2)
}

func NewSub(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return nil, err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return nil, err
		}

		return i1 - i2, nil
	}, c1, c2)
}

func NewOR(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		b1, err := callBool(c1, c)
		if err != nil {
			return nil, err
		}
		b2, err := callBool(c2, c)
		if err != nil {
			return nil, err
		}
		return b1 || b2, nil
	}, c1, c2)
}

func NewAND(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		b1, err := callBool(c1, c)
		if err != nil {
			return nil, err
		}
		b2, err := callBool(c2, c)
		if err != nil {
			return nil, err
		}
		return b1 && b2, nil
	}, c1, c2)
}

func NewEqual(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		v1, err := c.Call(c1)
		if err != nil {
			return nil, err
		}
		v2, err := c.Call(c2)
		if err != nil {
			return nil, err
		}

		return Equals(v1, v2), nil
	}, c1, c2)
}

func NewNotEqual(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		v1, err := c.Call(c1)
		if err != nil {
			return nil, err
		}
		v2, err := c.Call(c2)
		if err != nil {
			return nil, err
		}

		return !Equals(v1, v2), nil
	}, c1, c2)
}

func NewLt(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return nil, err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return nil, err
		}

		return i1 < i2, nil
	})
}

func NewGt(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return nil, err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return nil, err
		}

		return i1 > i2, nil
	})
}
