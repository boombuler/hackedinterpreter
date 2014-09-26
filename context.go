package main

import (
	"errors"
	"fmt"
	"github.com/boombuler/gold"
	"time"
)

const NoTimeout time.Duration = 0
const DefaultTimeout = 5 * time.Second

type Context struct {
	parentContext *Context
	variables     map[string]Value
	functions     [16]*Function

	result  Value
	err     error
	timeout *time.Timer
	ui      UIInterface
}

func (c *Context) SetInput(val Value) {
	c.variables["input"] = val
}

func NewContext(d time.Duration) *Context {
	c := new(Context)
	c.variables = make(map[string]Value)
	if d != 0 {
		c.timeout = time.NewTimer(d)
	}
	return c
}

func (c *Context) UI() UIInterface {
	for c.parentContext != nil {
		if c.ui != nil {
			return c.ui
		}
		c = c.parentContext
	}
	if c.ui == nil {
		c.ui = new(UIStub)
	}
	return c.ui
}

func (c *Context) newChildContext() *Context {
	return &Context{
		variables:     make(map[string]Value),
		parentContext: c,
	}
}

func (c *Context) forceExit() bool {
	if c.err != nil || c.result != nil {
		return true
	}
	if c.parentContext != nil {
		return c.parentContext.forceExit()
	} else if c.timeout != nil {
		select {
		case <-c.timeout.C:
			c.err = errors.New("timeout")
			c.result = nil
			return true
		default:
			return false
		}
	} else {
		return false
	}
}

func (c *Context) getFunctionIdx(sym gold.SymbolId) int {
	switch sym {
	case Sym_F1:
		return 0
	case Sym_F2:
		return 1
	case Sym_F3:
		return 2
	case Sym_F4:
		return 3
	case Sym_F5:
		return 4
	case Sym_F6:
		return 5
	case Sym_F7:
		return 6
	case Sym_F8:
		return 7
	case Sym_F9:
		return 8
	case Sym_F10:
		return 9
	case Sym_F11:
		return 10
	case Sym_F12:
		return 11
	case Sym_F13:
		return 12
	case Sym_F14:
		return 13
	case Sym_F15:
		return 14
	case Sym_F16:
		return 15
	default:
		return -1
	}
}

func (c *Context) setFunction(sym gold.SymbolId, f *Function) error {
	idx := c.getFunctionIdx(sym)
	if idx < 0 {
		return errors.New("invalid function name")
	}
	c.functions[idx] = f
	return nil
}

func (c *Context) getFunction(sym gold.SymbolId) (*Function, error) {
	idx := c.getFunctionIdx(sym)
	if idx < 0 {
		return nil, errors.New("invalid function name")
	}
	f := c.functions[idx]
	for f == nil && c.parentContext != nil {
		c = c.parentContext
		f = c.functions[idx]
	}
	if f != nil {
		return f, nil
	} else {
		return nil, fmt.Errorf("function f%v not defined", idx+1)
	}
}

func varNameFromToken(t *gold.Token) (string, error) {
	if t.Rule == Rule_VariableVariable || t.Rule == Rule_VariableInput {
		t = t.Tokens[0]
	}

	if t.Symbol == Sym_Variable {
		return t.Text, nil
	} else if t.Symbol == Sym_Input {
		return "input", nil
	} else {
		fmt.Println(t)
		return "", errors.New("variable expected")
	}

}

func ValueVariable(t *gold.Token, c *Context) (Value, error) {
	vn, err := varNameFromToken(t.Tokens[0])
	if err != nil {
		return nil, err
	}
	val, ok := c.variables[vn]
	if ok {
		return val, nil
	} else {
		return 0, nil
	}
}

func AssignstmntVariableEq(t *gold.Token, c *Context) (Value, error) {
	vn, err := varNameFromToken(t.Tokens[0])
	if err != nil {
		return nil, err
	}
	val, err := Exec(t.Tokens[2], c)
	if err != nil {
		return val, err
	}
	if val == nil {
		return nil, errors.New("Value expected")
	}
	c.variables[vn] = val
	return val, nil
}

func Return(t *gold.Token, c *Context) (Value, error) {
	c.result, c.err = Exec(t.Tokens[1], c)
	return c.result, c.err
}
