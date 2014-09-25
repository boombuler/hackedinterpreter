package main

import (
	"errors"
	"fmt"
	"github.com/boombuler/gold"
)

func invokeMethod(sender Value, method *gold.Token, params []Value, c *Context) (Value, error) {
	list, isLst := sender.(*List)
	if method.Symbol == Sym_Is_list {
		// only method that can be called on other things than lists
		if len(params) != 0 {
			return false, errors.New("no parameters expected")
		}
		return isLst, nil
	}
	if !isLst {
		return nil, fmt.Errorf("can not call '%v' on %v", method.Text, getTypeName(sender))
	}

	switch method.Symbol {
	case Sym_Length:
		if len(params) != 0 {
			return 0, errors.New("no parameters expected")
		}
		return len(list.content), nil
	case Sym_Sort:
		if len(params) != 0 {
			return 0, errors.New("no parameters expected")
		}
		return list, list.Sort(DefaultSort)
	case Sym_Copy:
		if len(params) != 0 {
			return 0, errors.New("no parameters expected")
		}
		return list.Copy(), nil
	case Sym_Push:
		if len(params) != 1 {
			return 0, errors.New("parameter missmatch. Syntax: push(Anything)")
		}
		list.Push(params[0])
		return list, nil
	case Sym_Pop:
		if len(params) != 0 {
			return 0, errors.New("no parameters expected")
		}
		return list.Pop()
	case Sym_Remove:
		if len(params) != 1 {
			return 0, errors.New("parameter missmatch. Syntax: remove(Integer)")
		}
		idx, ok := params[0].(int)
		if !ok {
			return 0, errors.New("parameter missmatch. Syntax: remove(Integer)")
		}
		return list, list.Remove(idx)
	case Sym_Fill:
		if len(params) != 1 {
			return 0, errors.New("parameter missmatch. Syntax: fill(Anything)")
		}
		for i, _ := range list.content {
			list.content[i] = params[0]
		}
		return list, nil
	case Sym_Insert:
		if len(params) != 2 {
			return 0, errors.New("parameter missmatch. Syntax: insert(Integer, Anything)")
		}
		idx, ok := params[0].(int)
		if !ok {
			return 0, errors.New("parameter missmatch. Syntax: insert(Integer, Anything)")
		}
		return list, list.Insert(idx, params[1])
	case Sym_Sort_with:
		if len(params) != 1 {
			return 0, errors.New("parameter missmatch. Syntax: sort_with(Function)")
		}
		function, ok := params[0].(*Function)
		if !ok {
			return 0, errors.New("parameter missmatch.  Syntax: sort_with(Function)")
		}
		return list, list.Sort(func(v1, v2 Value) (bool, error) {
			res, err := function.Call([]Value{v1, v2}, c)
			if err != nil {
				return false, err
			}
			bRes, ok := res.(bool)
			if !ok {
				return false, errors.New("sort function parameter must return Boolean")
			}
			return bRes, nil
		})
	case Sym_Map:
		if len(params) != 1 {
			return 0, errors.New("parameter missmatch. Syntax: map(Function)")
		}
		function, ok := params[0].(*Function)
		if !ok {
			return 0, errors.New("parameter missmatch.  Syntax: map(Function)")
		}

		for i := 0; i < len(list.content); i++ {
			newVal, err := function.Call([]Value{list.content[i]}, c)
			if err != nil {
				return list, err
			}
			list.content[i] = newVal
		}
		return list, nil
	}

	return nil, errors.New("Invalid method name")
}

func CallMethodParam(t *gold.Token, c *Context) (Value, error) {
	// <Method Call> ::= <Value> '.' <Method Name> '(' <ArrayValues> ')'
	sender, err := Exec(t.Tokens[0], c)
	if err != nil {
		return sender, err
	}
	method := t.Tokens[2]
	if method.Symbol == Sym_Methodname {
		method = t.Tokens[2].Tokens[0]
	}
	params, err := getListValues(t.Tokens[4], c)
	if err != nil {
		return nil, err
	}

	return invokeMethod(sender, method, params, c)
}

func CallMethodNoParam(t *gold.Token, c *Context) (Value, error) {
	// <Method Call> ::= <Value> '.' <Method Name>
	sender, err := Exec(t.Tokens[0], c)
	if err != nil {
		return sender, err
	}
	method := t.Tokens[2]
	if method.Symbol == Sym_Methodname {
		method = t.Tokens[2].Tokens[0]
	}
	return invokeMethod(sender, method, []Value{}, c)
}
