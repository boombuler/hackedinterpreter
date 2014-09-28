package runtime

import (
	"errors"
)

func MakeMul(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return Bool(false), err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return Bool(false), err
		}

		return Int(i1 * i2), nil
	})
}

func MakeDiv(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return Bool(false), err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return Bool(false), err
		}
		if i2 == 0 {
			return 0, errors.New("Division by zero!")
		}

		return Int(i1 / i2), nil
	})
}

func MakeAdd(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return Bool(false), err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return Bool(false), err
		}

		return Int(i1 + i2), nil
	})
}

func MakeSub(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return Bool(false), err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return Bool(false), err
		}

		return Int(i1 - i2), nil
	})
}

func MakeOR(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		b1, err := callBool(c1, c)
		if err != nil {
			return Bool(false), err
		}
		b2, err := callBool(c2, c)
		if err != nil {
			return Bool(false), err
		}
		return Bool(b1 || b2), nil
	})
}

func MakeAND(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		b1, err := callBool(c1, c)
		if err != nil {
			return Bool(false), err
		}
		b2, err := callBool(c2, c)
		if err != nil {
			return Bool(false), err
		}
		return Bool(b1 && b2), nil
	})
}

func MakeEqual(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		v1, err := c1.Call(c)
		if err != nil {
			return Bool(false), err
		}
		v2, err := c2.Call(c)
		if err != nil {
			return Bool(false), err
		}

		return Bool(compare(v1, v2)), nil
	})
}

func MakeNotEqual(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		v1, err := c1.Call(c)
		if err != nil {
			return Bool(false), err
		}
		v2, err := c2.Call(c)
		if err != nil {
			return Bool(false), err
		}

		return Bool(!compare(v1, v2)), nil
	})
}

func MakeLt(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return Bool(false), err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return Bool(false), err
		}

		return Bool(i1 < i2), nil
	})
}

func MakeGt(c1, c2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		i1, err := callInt(c1, c)
		if err != nil {
			return Bool(false), err
		}
		i2, err := callInt(c2, c)
		if err != nil {
			return Bool(false), err
		}

		return Bool(i1 > i2), nil
	})
}
