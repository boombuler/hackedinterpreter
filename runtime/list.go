package runtime

import (
	"errors"
)

type List struct {
	content []Value
}

func newList(initLen int64) *List {
	c := make([]Value, initLen)
	for i := int64(0); i < initLen; i++ {
		c[i] = Int(0)
	}
	return &List{
		content: c,
	}
}

func (l *List) Len() Int {
	return Int(len(l.content))
}

func MakeEmptyList() Callable {
	return callableFunc(func(c *Context) (Value, error) {
		return newList(0), nil
	})
}

func MakeListValues(values Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		valsV, err := values.Call(c)
		if err != nil {
			return nil, err
		}
		vals := valsV.([]Value)
		lst := newList(int64(len(vals)))
		for i, v := range vals {
			lst.content[i] = v
		}

		return lst, nil
	})
}

func MakeGetListItem(list, index Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		lVal, err := callList(list, c)
		if err != nil {
			return nil, err
		}
		idx, err := callInt(index, c)
		if err != nil {
			return nil, err
		}
		if idx < 0 || idx >= lVal.Len() {
			return nil, errors.New("Index out of range")
		}
		return lVal.content[idx], nil
	})
}

func MakeSetListItem(list, index, value Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		lVal, err := callList(list, c)
		if err != nil {
			return nil, err
		}
		idx, err := callInt(index, c)
		if err != nil {
			return nil, err
		}
		val, err := value.Call(c)
		if err != nil {
			return nil, err
		}
		if idx < 0 || idx >= lVal.Len() {
			return nil, errors.New("Index out of range")
		}
		lVal.content[idx] = val
		return val, nil
	})
}
