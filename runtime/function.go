package runtime

import (
	"errors"
	"fmt"
	"github.com/boombuler/hackedinterpreter/token"
	"math"
	"math/rand"
	"time"
)

type buildInFn func(params []Value, c *Context) (Value, error)

type buildInFnSig struct {
	ParamTypes []ValueType
	Func       buildInFn
}

var buildInFunctions map[string]buildInFnSig = map[string]buildInFnSig{
	"abs": {[]ValueType{INT}, func(params []Value, c *Context) (Value, error) {
		i := params[0].(int)
		if i < 0 {
			return int(-i), nil
		}
		return int(i), nil
	}},
	"mod": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		return params[0].(int) % params[1].(int), nil
	}},
	"min": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(int)
		i2 := params[1].(int)
		if i1 < i2 {
			return i1, nil
		}
		return i2, nil
	}},
	"max": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(int)
		i2 := params[1].(int)
		if i1 > i2 {
			return i1, nil
		}
		return i2, nil
	}},
	"pow": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(int)
		i2 := params[1].(int)

		return int(math.Pow(float64(i1), float64(i2))), nil
	}},
	"new_list": {[]ValueType{INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(int)
		if i1 < 0 {
			return nil, errors.New("parameter as to be > 0")
		}
		return newList(i1), nil
	}},
	"draw": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			ui.Draw(params[0].(int), params[1].(int))
			return DRAWINGSF, nil
		}
		return nil, errors.New("UI functions not available")
	}},
	"draw_text": {[]ValueType{INT, INT, ANY}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			ui.DrawText(params[0].(int), params[1].(int), ToString(params[2], false))
			return DRAWINGSF, nil
		}
		return nil, errors.New("UI functions not available")
	}},
	"width": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.Width()
		}
		return nil, errors.New("UI functions not available")
	}},
	"height": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.Height()
		}
		return nil, errors.New("UI functions not available")
	}},
	"left": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.Left()
		}
		return nil, errors.New("UI functions not available")
	}},
	"right": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.Right()
		}
		return nil, errors.New("UI functions not available")
	}},
	"up": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.Up()
		}
		return nil, errors.New("UI functions not available")
	}},
	"down": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.Down()
		}
		return nil, errors.New("UI functions not available")
	}},
	"a_btn": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.BtnA()
		}
		return nil, errors.New("UI functions not available")
	}},
	"b_btn": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if ui := c.ui(); ui != nil {
			return ui.BtnB()
		}
		return nil, errors.New("UI functions not available")
	}},
	"random": {[]ValueType{INT}, func(params []Value, c *Context) (Value, error) {
		n := params[0].(int)
		if n <= 0 {
			return 0, errors.New("random parameter must be > 0")
		}

		return rand.Intn(n), nil
	}},
	"time": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		d := time.Now().Sub(startTime)
		return int(d.Seconds() * 1000), nil
	}},
}

var startTime time.Time

func init() {
	startTime = time.Now()
	rand.Seed(time.Now().UnixNano())
}

func getCustFnIdx(name string) int {
	switch name {
	case "f1":
		return 0
	case "f2":
		return 1
	case "f3":
		return 2
	case "f4":
		return 3
	case "f5":
		return 4
	case "f6":
		return 5
	case "f7":
		return 6
	case "f8":
		return 7
	case "f9":
		return 8
	case "f10":
		return 9
	case "f11":
		return 10
	case "f12":
		return 11
	case "f13":
		return 12
	case "f14":
		return 13
	case "f15":
		return 14
	case "f16":
		return 15
	}
	return -1
}

func (f *Function) Call(values []Value, c *Context) (Value, error) {
	cc := c.newChildContext()

	if len(f.paramNames) != len(values) {
		return nil, errors.New("parameter count missmatch")
	}
	for i, val := range values {
		cc.SetVariable(f.paramNames[i], val)
	}
	return cc.Call(f.body)
}

func NewCallFunction(fn string, values Callable, p *token.Token) (Callable, error) {
	var sig *buildInFnSig = nil

	psig, ok := buildInFunctions[fn]
	if !ok {
		return nil, errors.New("unknown function: " + fn)
	}
	sig = &psig

	return newCallable(p, func(c *Context) (Value, error) {
		var vals = []Value{}
		if values != nil {
			valsV, err := c.Call(values)
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

		return sig.Func(vals, c)
	}, values), nil
}

func NewCustFuncDev(name string, params []string, code Callable, p *token.Token) (Callable, error) {
	idx := getCustFnIdx(name)

	if idx < 0 {
		return nil, errors.New("Invalid function name: " + name)
	}
	funcDef := &Function{
		paramNames: params,
		body:       code,
	}

	return newCallable(p, func(c *Context) (Value, error) {
		c.functions[idx] = funcDef
		return funcDef, nil
	}, code), nil
}

type Function struct {
	paramNames []string
	body       Callable
}

func NewAddToParamDef(params []string, value string) []string {
	return append(params, value)
}

func NewParamDef(value string) []string {
	return []string{value}
}

func NewLambdaDef(params []string, code Callable, p *token.Token) Callable {
	lambda := &Function{
		paramNames: params,
		body:       code,
	}
	return newCallable(p, func(c *Context) (Value, error) {
		return lambda, nil
	}, code)
}

func NewGetCustomFunction(cfName string, p *token.Token) (Callable, error) {
	fn_idx := getCustFnIdx(cfName)
	if fn_idx < 0 {
		return nil, fmt.Errorf("Unknown custom function %v", cfName)
	}
	return newCallable(p, func(c *Context) (Value, error) {
		custFn := c.functions[fn_idx]
		cc := c.parentContext
		for custFn == nil && cc != nil {
			custFn = cc.functions[fn_idx]
			cc = cc.parentContext
		}
		if custFn == nil {
			return nil, fmt.Errorf("custom function %v is not defined", cfName)
		}
		return custFn, nil
	}), nil
}

func NewCallLambda(fn, values Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		ff, err := c.Call(fn)
		if err != nil {
			return ff, err
		}
		f, ok := ff.(*Function)
		if !ok {
			return ff, errors.New("Function expected!")
		}

		valsV, err := c.Call(values)
		if err != nil {
			return nil, err
		}
		vals := valsV.([]Value)

		if len(vals) != len(f.paramNames) {
			return ff, errors.New("Parameter count missmatch")
		}

		cc := c.newChildContext()
		for i, val := range vals {
			cc.SetVariable(f.paramNames[i], val)
		}
		return cc.Call(f.body)
	}, values)
}
