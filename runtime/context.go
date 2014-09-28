package runtime

type Context struct {
	parentContext *Context
	variables     map[string]Value
	functions     [16]*Function

	ui UIInterface
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

func NewContext() *Context {
	return &Context{
		variables: make(map[string]Value),
	}
}

func (c *Context) newChildContext() *Context {
	return &Context{
		variables:     make(map[string]Value),
		parentContext: c,
	}
}

func MakeGetVariable(varName []byte) Callable {
	vn := string(varName)
	return callableFunc(func(c *Context) (Value, error) {
		val, ok := c.variables[vn]
		if ok {
			return val, nil
		}
		return Int(0), nil
	})
}

func MakeSetVariable(varName []byte, value Callable) Callable {
	vn := string(varName)
	return callableFunc(func(c *Context) (Value, error) {
		val, err := value.Call(c)
		if err != nil {
			return nil, err
		}
		c.variables[vn] = val
		return val, nil
	})
}
