package runtime

import (
	"../token"
	"errors"
	"fmt"
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
		v1, err := c.Call(c1)
		if err != nil {
			return v1, err
		}
		v2, err := c.Call(c2)
		if err != nil {
			return v1, err
		}
		t1 := GetType(v1)
		t2 := GetType(v2)

		if t1 == STRING || t2 == STRING {
			return ToString(v1, false) + ToString(v2, false), nil
		} else if t1 == INT && t2 == INT {
			return v1.(int) + v2.(int), nil
		} else {
			return nil, fmt.Errorf("Can not add %v and %v", t1, t2)
		}
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
		v1, err := c.Call(c1)
		if err != nil {
			return v1, err
		}
		typ := GetType(v1)
		if typ == STRING {
			s1 := v1.(string)
			s2, err := callString(c2, c)
			if err != nil {
				return s2, err
			}
			return s1 < s2, nil
		} else if typ == INT {
			i1 := v1.(int)
			i2, err := callInt(c2, c)
			if err != nil {
				return nil, err
			}

			return i1 < i2, nil
		} else {
			return nil, errors.New("Integer or String expected")
		}
	})
}

func NewGt(c1, c2 Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		v1, err := c.Call(c1)
		if err != nil {
			return v1, err
		}
		typ := GetType(v1)
		if typ == STRING {
			s1 := v1.(string)
			s2, err := callString(c2, c)
			if err != nil {
				return s2, err
			}
			return s1 > s2, nil
		} else if typ == INT {
			i1 := v1.(int)
			i2, err := callInt(c2, c)
			if err != nil {
				return nil, err
			}

			return i1 > i2, nil
		} else {
			return nil, errors.New("Integer or String expected")
		}
	})
}
