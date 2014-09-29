package runtime

import (
	"errors"
	"fmt"
)

type methodFn func(sender Value, params []Value, c *Context) (Value, error)

type methodSig struct {
	SenderType ValueType
	ParamTypes []ValueType
	Func       methodFn
}

var methods map[string]methodSig = map[string]methodSig{
	"is_list": {ANY, []ValueType{}, func(sender Value, params []Value, c *Context) (Value, error) {
		return Bool(GetType(sender) == LIST), nil
	}},
	"length": {LIST, []ValueType{}, func(sender Value, params []Value, c *Context) (Value, error) {
		return (sender.(*List)).Len(), nil
	}},
	"push": {LIST, []ValueType{ANY}, func(sender Value, params []Value, c *Context) (Value, error) {
		(sender.(*List)).Push(params[0])
		return sender, nil
	}},
	"pop": {LIST, []ValueType{}, func(sender Value, params []Value, c *Context) (Value, error) {
		return (sender.(*List)).Pop()
	}},
	"insert": {LIST, []ValueType{INT, ANY}, func(sender Value, params []Value, c *Context) (Value, error) {
		return sender, (sender.(*List)).Insert(int(params[0].(Int)), params[1])
	}},
	"remove": {LIST, []ValueType{INT}, func(sender Value, params []Value, c *Context) (Value, error) {
		return sender, (sender.(*List)).Remove(int(params[0].(Int)))
	}},
	"sort": {LIST, []ValueType{}, func(sender Value, params []Value, c *Context) (Value, error) {
		list := sender.(*List)
		return list, list.sort(defaultSort)
	}},
	"sort_with": {LIST, []ValueType{FUNC}, func(sender Value, params []Value, c *Context) (Value, error) {
		list := sender.(*List)
		sortFn := params[0].(*Function)
		return list, list.sort(func(v1, v2 Value) (bool, error) {
			res, err := sortFn.Call([]Value{v1, v2}, c)
			if err != nil {
				return false, err
			}
			bRes, ok := res.(Bool)
			if !ok {
				return false, errors.New("sort function parameter must return Boolean")
			}
			return bool(bRes), nil
		})
	}},
	"map": {LIST, []ValueType{FUNC}, func(sender Value, params []Value, c *Context) (Value, error) {
		list := sender.(*List)
		mapFn := params[0].(*Function)

		for i := 0; i < len(list.content); i++ {
			newVal, err := mapFn.Call([]Value{list.content[i]}, c)
			if err != nil {
				return list, err
			}
			list.content[i] = newVal
		}
		return list, nil
	}},
	"copy": {LIST, []ValueType{}, func(sender Value, params []Value, c *Context) (Value, error) {
		return (sender.(*List)).Copy(), nil
	}},
	"fill": {LIST, []ValueType{ANY}, func(sender Value, params []Value, c *Context) (Value, error) {
		list := sender.(*List)
		for i, _ := range list.content {
			list.content[i] = params[0]
		}
		return list, nil
	}},
}

func NewMethodCall(sender Callable, methodName string, values Callable) (Callable, error) {
	sig, ok := methods[methodName]
	if !ok {
		return nil, fmt.Errorf("Unknown method: %v", methodName)
	}

	return callableFunc(func(c *Context) (Value, error) {
		send, err := sender.Call(c)
		if err != nil {
			return send, err
		}

		if st := GetType(send); st != sig.SenderType && sig.SenderType != ANY {
			return nil, fmt.Errorf("Can not call %v on %v", methodName, GetType(send))
		}

		var vals = []Value{}
		if values != nil {
			valsV, err := values.Call(c)
			if err != nil {
				return nil, err
			}
			vals = valsV.([]Value)
		}

		if len(vals) != len(sig.ParamTypes) {
			return nil, errors.New("parameter count missmatch")
		}

		for i := 0; i < len(vals); i++ {
			pt := GetType(vals[i])
			wt := sig.ParamTypes[i]
			if pt != wt && wt != ANY {
				return nil, fmt.Errorf("wrong parameter type. got %v but expected %v", pt, wt)
			}
		}

		return sig.Func(send, vals, c)

	}), nil
}
