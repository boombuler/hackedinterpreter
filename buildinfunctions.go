package main

import (
	"errors"
	"fmt"
	"github.com/boombuler/gold"
	"math"
	"math/rand"
	"time"
)

type BuildInFn func(params []Value, c *Context) (Value, error)

type BuildInFnSig struct {
	ParamTypes []ValueType
	Func       BuildInFn
}

var BuildInFunctions map[gold.RuleId]BuildInFnSig = map[gold.RuleId]BuildInFnSig{
	Rule_FunctionnameAbs:       {[]ValueType{INT}, Abs},
	Rule_FunctionnameMod:       {[]ValueType{INT, INT}, Mod},
	Rule_FunctionnameMin:       {[]ValueType{INT, INT}, Min},
	Rule_FunctionnameMax:       {[]ValueType{INT, INT}, Max},
	Rule_FunctionnameRandom:    {[]ValueType{INT}, Random},
	Rule_FunctionnameTime:      {[]ValueType{}, Time},
	Rule_FunctionnamePow:       {[]ValueType{INT, INT}, Pow},
	Rule_FunctionnameNew_list:  {[]ValueType{INT}, FnNewList},
	Rule_FunctionnameB_btn:     {[]ValueType{}, BtnB},
	Rule_FunctionnameA_btn:     {[]ValueType{}, BtnA},
	Rule_FunctionnameDown:      {[]ValueType{}, BtnDown},
	Rule_FunctionnameUp:        {[]ValueType{}, BtnUp},
	Rule_FunctionnameLeft:      {[]ValueType{}, BtnLeft},
	Rule_FunctionnameRight:     {[]ValueType{}, BtnRight},
	Rule_FunctionnameWidth:     {[]ValueType{}, Width},
	Rule_FunctionnameHeight:    {[]ValueType{}, Height},
	Rule_FunctionnameDraw_text: {[]ValueType{INT, INT, ANY}, DrawText},
	Rule_FunctionnameDraw:      {[]ValueType{INT, INT}, Draw},
}

func Draw(params []Value, c *Context) (Value, error) {
	return DRAWINGSF, c.UI().Draw(params[0].(int), params[1].(int))
}
func DrawText(params []Value, c *Context) (Value, error) {
	return DRAWINGSF, c.UI().DrawText(params[0].(int), params[1].(int), ToString(params[2]))
}
func Height(params []Value, c *Context) (Value, error) {
	return c.UI().Height()
}
func Width(params []Value, c *Context) (Value, error) {
	return c.UI().Width()
}
func BtnDown(params []Value, c *Context) (Value, error) {
	return c.UI().Down()
}
func BtnUp(params []Value, c *Context) (Value, error) {
	return c.UI().Up()
}
func BtnLeft(params []Value, c *Context) (Value, error) {
	return c.UI().Left()
}
func BtnRight(params []Value, c *Context) (Value, error) {
	return c.UI().Right()
}

func BtnA(params []Value, c *Context) (Value, error) {
	return c.UI().BtnA()
}
func BtnB(params []Value, c *Context) (Value, error) {
	return c.UI().BtnB()
}
func FnNewList(params []Value, c *Context) (Value, error) {
	length := params[0].(int)
	if length < 0 {
		return nil, errors.New("list size must be >= 0")
	}

	l := &List{
		content: make([]Value, length, length),
	}
	for i := 0; i < length; i++ {
		l.content[i] = 0
	}
	return l, nil
}

func Pow(params []Value, c *Context) (Value, error) {
	a := params[0].(int)
	b := params[1].(int)
	return int(math.Pow(float64(a), float64(b))), nil
}

func Time(params []Value, c *Context) (Value, error) {
	d := time.Now().Sub(StartTime)
	return int(int(d.Seconds() * 1000)), nil
}

func Random(params []Value, c *Context) (Value, error) {
	n := params[0].(int)
	if n <= 0 {
		return 0, errors.New("random parameter must be > 0")
	}

	return int(rand.Int63n(int64(n))), nil
}
func Mod(params []Value, c *Context) (Value, error) {
	v1 := params[0].(int)
	v2 := params[1].(int)
	return v1 % v2, nil
}

func Min(params []Value, c *Context) (Value, error) {
	v1 := params[0].(int)
	v2 := params[1].(int)
	if v1 < v2 {
		return v1, nil
	}
	return v2, nil
}

func Max(params []Value, c *Context) (Value, error) {
	v1 := params[0].(int)
	v2 := params[1].(int)
	if v1 < v2 {
		return v1, nil
	}
	return v2, nil
}

func Abs(params []Value, c *Context) (Value, error) {
	val := params[0].(int)
	if val < 0 {
		return -val, nil
	}
	return val, nil
}

func CallBuildInFn(t *gold.Token, c *Context) (Value, error) {
	// <Function Call> ::= <Function Name> '(' <Fn Parameters> ')'
	params, err := getListValues(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}
	sig, ok := BuildInFunctions[t.Tokens[0].Rule]
	if !ok {
		return nil, errors.New("function not implemented!")
	}

	if len(params) != len(sig.ParamTypes) {
		return nil, errors.New("parameter count missmatch")
	}

	for i := 0; i < len(params); i++ {
		pt := getTypeName(params[i])
		wt := sig.ParamTypes[i]
		if pt != wt && wt != ANY {
			return nil, fmt.Errorf("wrong parameter type. got %v but expected %v", pt, wt)
		}
	}

	return sig.Func(params, c)
}
