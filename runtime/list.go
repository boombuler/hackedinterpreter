package runtime

import (
	"../token"
	"errors"
	"fmt"
	"sort"
)

type List struct {
	content []Value
}

func newList(initLen int) *List {
	c := make([]Value, initLen)
	for i := 0; i < initLen; i++ {
		c[i] = 0
	}
	return &List{
		content: c,
	}
}

func (l *List) sort(less listSorterCompare) error {
	helper := &listSortHelper{l, less, nil}
	sort.Sort(helper)
	return helper.err
}

func defaultSort(a, b Value) (bool, error) {
	intA, ok := a.(int)
	if ok {
		intB, ok := b.(int)
		if ok {
			return intA < intB, nil
		}
		return false, fmt.Errorf("can not compare '%v' and '%v'", GetType(a), GetType(b))
	}
	boolA, ok := a.(bool)
	if ok {
		_, ok := b.(bool)
		if ok {
			return !boolA, nil
		}
		return false, fmt.Errorf("can not compare '%v' and '%v'", GetType(a), GetType(b))
	}
	return false, fmt.Errorf("can not compare '%v' and '%v'. Try to use 'sort_with'", GetType(a), GetType(b))
}

type listSorterCompare func(a, b Value) (bool, error)

type listSortHelper struct {
	*List
	IsLess listSorterCompare
	err    error
}

func (s *listSortHelper) Len() int {
	return len(s.content)
}
func (s *listSortHelper) Swap(i, j int) {
	s.content[i], s.content[j] = s.content[j], s.content[i]
}
func (s *listSortHelper) Less(i, j int) bool {
	res, err := s.IsLess(s.content[i], s.content[j])
	if err != nil {
		s.err = err
	}
	return res
}

func (l *List) Copy() *List {
	return &List{
		content: append([]Value{}, l.content...),
	}
}

func (l *List) Len() int {
	return len(l.content)
}

func (l *List) Push(v Value) {
	l.content = append(l.content, v)
}

func (l *List) Pop() (Value, error) {
	if len(l.content) == 0 {
		return nil, errors.New("Can not pop from empty list")
	}
	ll := len(l.content)
	val := l.content[ll-1]
	l.content = l.content[0 : ll-1]
	return val, nil
}

func (l *List) Insert(idx int, v Value) error {
	if idx < 0 || idx >= len(l.content) {
		return fmt.Errorf("Index out of bounds [%v]", idx)
	}
	l.content = append(l.content[:idx], append([]Value{v}, l.content[idx:]...)...)
	return nil
}

func (l *List) Remove(idx int) error {
	if idx < 0 || idx >= len(l.content) {
		return fmt.Errorf("Index out of bounds [%v]", idx)
	}
	copy(l.content[idx:], l.content[idx+1:])
	l.content[len(l.content)-1] = nil
	l.content = l.content[:len(l.content)-1]
	return nil
}

func NewEmptyList(p *token.Pos) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		return newList(0), nil
	})
}

func NewListValues(values *Callable, p *token.Pos) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		valsV, err := values.Call(c)
		if err != nil {
			return nil, err
		}
		vals := valsV.([]Value)
		lst := newList(len(vals))
		for i, v := range vals {
			lst.content[i] = v
		}

		return lst, nil
	})
}

func NewGetListItem(list, index *Callable, p *token.Pos) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		lVal, err := callList(list, c)
		if err != nil {
			return nil, err
		}
		idx, err := callInt(index, c)
		if err != nil {
			return nil, err
		}
		if idx < 0 || idx >= lVal.Len() {
			return nil, errors.New("Index out of range (" + ToString(idx) + ") -> " + ToString(lVal))
		}
		return lVal.content[idx], nil
	})
}

func NewSetListItem(list, index, value *Callable, p *token.Pos) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
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
