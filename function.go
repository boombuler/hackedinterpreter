package main

import (
	"errors"
	"fmt"
	"github.com/boombuler/gold"
)

type Function struct {
	paramNames []string
	body       *gold.Token
}

func (f *Function) Call(parameters []Value, c *Context) (Value, error) {
	cc := c.newChildContext()
	if len(parameters) != len(f.paramNames) {
		return nil, errors.New("parameter count missmatch")
	}
	for i, p := range f.paramNames {
		cc.variables[p] = parameters[i]
	}
	return Exec(f.body, cc)
}

func makeFunction(pdef, body *gold.Token) (*Function, error) {
	params := []string{}

	for {
		if pdef.Rule == Rule_FnparametersdefComma {
			vn, err := varNameFromToken(pdef.Tokens[2])
			if err != nil {
				return nil, err
			}

			params = append([]string{vn}, params...)
			pdef = pdef.Tokens[0]
		} else {
			vn, err := varNameFromToken(pdef)
			if err != nil {
				return nil, err
			}
			params = append([]string{vn}, params...)
			break
		}
	}
	return &Function{
		body:       body,
		paramNames: params,
	}, nil
}

func DefineFunction(t *gold.Token, c *Context) (Value, error) {
	// <Cust Fn Def> ::= function <Cust Fn Name> ':' <Fn Parameters Def> <CodeBlock>
	fnName := t.Tokens[1].Tokens[0].Symbol

	f, err := makeFunction(t.Tokens[3], t.Tokens[4])
	if err != nil {
		return nil, err
	}
	return f, c.setFunction(fnName, f)
}

func DefineLambda(t *gold.Token, c *Context) (Value, error) {
	// <Lambda Fn Def> ::= function <Fn Parameters Def> '->' <Statement>
	return makeFunction(t.Tokens[1], t.Tokens[3])
}

func GetFunctionCallParams(pdef *gold.Token, c *Context) ([]Value, error) {
	params := []Value{}

	for {
		if pdef.Rule == Rule_FnparametersComma {
			val, err := Exec(pdef.Tokens[2], c)
			if err != nil {
				return nil, err
			}
			params = append([]Value{val}, params...)
			pdef = pdef.Tokens[0]
		} else {
			val, err := Exec(pdef, c)
			if err != nil {
				return nil, err
			}
			params = append([]Value{val}, params...)
			break
		}
	}
	return params, nil
}

func CallFunction(t *gold.Token, c *Context) (Value, error) {
	// <Cust Fn Call> ::= <Cust Fn Name> '(' <Cust Fn Parameters> ')'
	fn, err := c.getFunction(t.Tokens[0].Tokens[0].Symbol)
	if err != nil {
		return nil, err
	}
	params, err := GetFunctionCallParams(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}
	return fn.Call(params, c)
}

func CallLambda(t *gold.Token, c *Context) (Value, error) {
	// <Lambda Fn Call> ::= <Lambda Target> '(' <Fn Parameters> ')'
	fnVal, err := Exec(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	fn, ok := fnVal.(*Function)
	if !ok {
		return nil, fmt.Errorf("function expected but got %v", getTypeName(fnVal))
	}

	params, err := GetFunctionCallParams(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}
	return fn.Call(params, c)
}
func GetFunction(t *gold.Token, c *Context) (Value, error) {
	return c.getFunction(t.Tokens[0].Symbol)
}
