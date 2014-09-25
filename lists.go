package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/boombuler/gold"
	"sort"
)

type List struct {
	content []Value
}

func getListValues(t *gold.Token, c *Context) ([]Value, error) {
	switch t.Rule {
	case Rule_ArrayvaluesComma:
		lst, err := getListValues(t.Tokens[0], c)
		if err != nil {
			return nil, err
		}
		entr, err := getListValues(t.Tokens[2], c)
		if err != nil {
			return nil, err
		}
		return append(lst, entr...), nil
	case Rule_Arrayvalues:
		entr, err := Exec(t.Tokens[0], c)
		if err != nil {
			return nil, err
		}
		return []Value{entr}, nil
	case Rule_Arrayvalues2:
		return []Value{}, nil
	}

	entr, err := Exec(t, c)
	if err != nil {
		return nil, err
	}
	return []Value{entr}, nil
}

func ConstList(t *gold.Token, c *Context) (Value, error) {
	values, err := getListValues(t.Tokens[1], c)
	if err != nil {
		return nil, err
	}
	return &List{content: values}, nil
}

func newList() *List {
	return &List{make([]Value, 0, 0)}
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

func ListGetIndex(t *gold.Token, c *Context) (Value, error) {
	lst, err := ExecWantList(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	idx, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	if idx < 0 || idx >= len(lst.content) {
		return nil, fmt.Errorf("Index out of bounds [%v]", idx)
	}
	return lst.content[idx], nil
}

func (l *List) Copy() *List {
	return &List{
		content: append([]Value{}, l.content...),
	}
}

func (l *List) String() string {
	buf := &bytes.Buffer{}
	buf.WriteRune('[')
	for i, val := range l.content {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%v", val))
	}

	buf.WriteRune(']')

	return buf.String()
}

func (l *List) Sort(less listSorterCompare) error {
	helper := &listSortHelper{l, less, nil}
	sort.Sort(helper)
	return helper.err
}

func DefaultSort(a, b Value) (bool, error) {
	intA, ok := a.(int)
	if ok {
		intB, ok := b.(int)
		if ok {
			return intA < intB, nil
		}
		return false, fmt.Errorf("can not compare '%v' and '%v'", getTypeName(a), getTypeName(b))
	}
	boolA, ok := a.(bool)
	if ok {
		_, ok := b.(bool)
		if ok {
			return !boolA, nil
		}
		return false, fmt.Errorf("can not compare '%v' and '%v'", getTypeName(a), getTypeName(b))
	}
	return false, fmt.Errorf("can not compare '%v' and '%v'. Try to use 'sort_with'", getTypeName(a), getTypeName(b))
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
