package runtime

import (
	"../token"
	"errors"
)

type Assignable interface {
	Callable
	Set(value Value, c *Context) error
}

type variable struct {
	*callable
	p  *token.Token
	vn string
}

func NewVariable(varName string, p *token.Token) Assignable {
	get := newCallable(p, func(c *Context) (Value, error) {
		val, ok := c.variables[varName]
		if ok {
			return val, nil
		}
		return int(0), nil
	})
	return &variable{get, p, varName}
}

func (v *variable) Set(value Value, c *Context) error {
	c.variables[v.vn] = value
	return nil
}

type listItem struct {
	*callable
	list  Callable
	index Callable
	p     *token.Token
}

func NewListItem(lst, idx Callable, p *token.Token) Assignable {
	get := newCallable(p, func(c *Context) (Value, error) {
		idx, err := callInt(idx, c)
		if err != nil {
			return nil, err
		}
		valList, err := c.Call(lst)
		if err != nil {
			return nil, err
		}
		listType := GetType(valList)
		if listType == LIST {
			lVal := valList.(*List)
			if err != nil {
				return nil, err
			}

			if idx < 0 || idx >= lVal.Len() {
				return nil, errors.New("Index out of range (" + ToString(idx, true) + ") -> " + ToString(lVal, true))
			}
			return lVal.content[idx], nil
		} else if listType == STRING {
			sVal := valList.(string)
			if !strIndexInRange(sVal, idx) {
				return nil, errors.New("Index out of range (" + ToString(idx, true) + ") -> " + ToString(sVal, true))
			}
			return strGetRuneByIdx(sVal, idx), nil
		} else {
			return nil, errors.New("Expected List or String")
		}
	}, lst, idx)
	return &listItem{get, lst, idx, p}
}

func (li *listItem) Set(value Value, c *Context) error {
	lVal, err := callList(li.list, c)
	if err != nil {
		return err
	}
	idx, err := callInt(li.index, c)
	if err != nil {
		return err
	}
	if idx < 0 || idx >= lVal.Len() {
		return errors.New("Index out of range")
	}
	lVal.content[idx] = value
	return nil
}

func NewAssign(a Assignable, value Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		v, err := c.Call(value)
		if err != nil {
			return v, err
		}
		return v, a.Set(v, c)
	}, a, value)
}
