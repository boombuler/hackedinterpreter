package runtime

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type BuildInFn func(params []Value, c *Context) (Value, error)

type BuildInFnSig struct {
	ParamTypes []ValueType
	Func       BuildInFn
}

var BuildInFunctions map[string]BuildInFnSig = map[string]BuildInFnSig{
	"abs": {[]ValueType{INT}, func(params []Value, c *Context) (Value, error) {
		i := params[0].(Int)
		if i < 0 {
			return Int(-i), nil
		}
		return Int(i), nil
	}},
	"mod": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		return Int(params[0].(Int) % params[1].(Int)), nil
	}},
	"min": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(Int)
		i2 := params[1].(Int)
		if i1 < i2 {
			return i1, nil
		}
		return i2, nil
	}},
	"max": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(Int)
		i2 := params[1].(Int)
		if i1 > i2 {
			return i1, nil
		}
		return i2, nil
	}},
	"pow": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(Int)
		i2 := params[1].(Int)

		return Int(int64(math.Pow(float64(i1), float64(i2)))), nil
	}},
	"new_list": {[]ValueType{INT}, func(params []Value, c *Context) (Value, error) {
		i1 := params[0].(Int)
		if i1 < 0 {
			return nil, errors.New("parameter as to be > 0")
		}
		return newList(int64(i1)), nil
	}},
	"draw": {[]ValueType{INT, INT}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			c.ui.Draw(int(params[0].(Int)), int(params[1].(Int)))
			return DRAWINGSF, nil
		}
		return nil, errors.New("UI functions not available")
	}},
	"draw_text": {[]ValueType{INT, INT, ANY}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			c.ui.DrawText(int(params[0].(Int)), int(params[1].(Int)), ToString(params[2]))
			return DRAWINGSF, nil
		}
		return nil, errors.New("UI functions not available")
	}},
	"width": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			w, err := c.ui.Width()
			return Int(w), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"height": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			h, err := c.ui.Height()
			return Int(h), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"left": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			l, err := c.ui.Left()
			return Bool(l), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"right": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			r, err := c.ui.Right()
			return Bool(r), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"up": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			u, err := c.ui.Up()
			return Bool(u), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"down": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			d, err := c.ui.Down()
			return Bool(d), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"a_btn": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			a, err := c.ui.BtnA()
			return Bool(a), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"b_btn": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		if c.ui != nil {
			b, err := c.ui.BtnB()
			return Bool(b), err
		}
		return nil, errors.New("UI functions not available")
	}},
	"random": {[]ValueType{INT}, func(params []Value, c *Context) (Value, error) {
		n := params[0].(Int)
		if n <= 0 {
			return 0, errors.New("random parameter must be > 0")
		}

		return Int(rand.Int63n(int64(n))), nil
	}},
	"time": {[]ValueType{}, func(params []Value, c *Context) (Value, error) {
		d := time.Now().Sub(startTime)
		return Int(int64(d.Seconds() * 1000)), nil
	}},
}

var startTime time.Time

func init() {
	startTime = time.Now()
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

func callCustomFn(idx int, values []Value, c *Context) (Value, error) {
	for c.functions[idx] == nil {
		c = c.parentContext
		if c == nil {
			return nil, fmt.Errorf("function f%v not defined", idx+1)
		}
	}
	f := c.functions[idx]
	cc := c.newChildContext()

	if len(f.paramNames) != len(values) {
		return nil, errors.New("parameter count missmatch")
	}
	for i, val := range values {
		cc.variables[f.paramNames[i]] = val
	}

	return f.body.Call(cc)
}

func CallFunction(funcName []byte, values Callable) (Callable, error) {
	fn := string(funcName)
	cf_idx := getCustFnIdx(fn)

	var sig *BuildInFnSig = nil

	if cf_idx < 0 {
		ssig, ok := BuildInFunctions[fn]
		if !ok {
			return nil, errors.New("unknown function: " + string(funcName))
		}
		sig = &ssig
	}

	return callableFunc(func(c *Context) (Value, error) {
		var vals = []Value{}
		if values != nil {
			valsV, err := values.Call(c)
			if err != nil {
				return nil, err
			}
			vals = valsV.([]Value)
		}

		if cf_idx >= 0 {
			return callCustomFn(cf_idx, vals, c)
		}

		if len(vals) != len(sig.ParamTypes) {
			return nil, errors.New("parameter count missmatch")
		}

		for i := 0; i < len(vals); i++ {
			pt := getType(vals[i])
			wt := sig.ParamTypes[i]
			if pt != wt && wt != ANY {
				return nil, fmt.Errorf("wrong parameter type. got %v but expected %v", pt, wt)
			}
		}

		return sig.Func(vals, c)
	}), nil
}

func MakeCustFuncDev(funcName []byte, params []string, code Callable) (Callable, error) {
	name := string(funcName)
	idx := getCustFnIdx(name)

	if idx < 0 {
		return nil, errors.New("Invalid function name: " + name)
	}
	funcDef := &Function{
		paramNames: params,
		body:       code,
	}

	return callableFunc(func(c *Context) (Value, error) {
		c.functions[idx] = funcDef
		return funcDef, nil
	}), nil
}

type Function struct {
	paramNames []string
	body       Callable
}

func AddToParamDef(params []string, value []byte) []string {
	return append(params, string(value))
}

func NewParamDef(value []byte) []string {
	return []string{string(value)}
}

func MakeLambda(params []string, code Callable) Callable {
	lambda := &Function{
		paramNames: params,
		body:       code,
	}
	return callableFunc(func(c *Context) (Value, error) {
		return lambda, nil
	})
}

func MakeCallLambda(fn, values Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		ff, err := fn.Call(c)
		if err != nil {
			return ff, err
		}
		f, ok := ff.(*Function)
		if !ok {
			return ff, errors.New("Function expected!")
		}

		valsV, err := values.Call(c)
		if err != nil {
			return nil, err
		}
		vals := valsV.([]Value)

		if len(vals) != len(f.paramNames) {
			return ff, errors.New("Parameter count missmatch")
		}

		cc := c.newChildContext()
		for i, val := range vals {
			cc.variables[f.paramNames[i]] = val
		}
		return f.body.Call(cc)
	})
}
