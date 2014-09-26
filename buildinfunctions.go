package main

import (
	"errors"
	"fmt"
	"github.com/boombuler/gold"
	"math"
	"math/rand"
	"time"
)

type BuildInFn func(params []Value) (Value, error)

type BuildInFnSig struct {
	ParamTypes []ValueType
	Func       BuildInFn
}

var BuildInFunctions map[gold.RuleId]BuildInFnSig = map[gold.RuleId]BuildInFnSig{
	Rule_FunctionnameAbs:      {[]ValueType{INT}, Abs},
	Rule_FunctionnameMod:      {[]ValueType{INT, INT}, Mod},
	Rule_FunctionnameMin:      {[]ValueType{INT, INT}, Min},
	Rule_FunctionnameMax:      {[]ValueType{INT, INT}, Max},
	Rule_FunctionnameRandom:   {[]ValueType{INT}, Random},
	Rule_FunctionnameTime:     {[]ValueType{}, Time},
	Rule_FunctionnamePow:      {[]ValueType{INT, INT}, Pow},
	Rule_FunctionnameNew_list: {[]ValueType{INT}, FnNewList},
	/*
	   Rule_FunctionnameB_btn
	   Rule_FunctionnameA_btn
	   Rule_FunctionnameDown
	   Rule_FunctionnameUp
	   Rule_FunctionnameRight
	   Rule_FunctionnameLeft
	   Rule_FunctionnameHeight
	   Rule_FunctionnameWidth
	   Rule_FunctionnameDraw_text
	   Rule_FunctionnameDraw
	*/
}

func FnNewList(params []Value) (Value, error) {
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

func Pow(params []Value) (Value, error) {
	a := params[0].(int)
	b := params[1].(int)
	return int(math.Pow(float64(a), float64(b))), nil
}

func Time(params []Value) (Value, error) {
	d := time.Now().Sub(StartTime)
	return int(int(d.Seconds() * 1000)), nil
}

func Random(params []Value) (Value, error) {
	n := params[0].(int)
	if n <= 0 {
		return 0, errors.New("random parameter must be > 0")
	}

	return int(rand.Int63n(int64(n))), nil
}
func Mod(params []Value) (Value, error) {
	v1 := params[0].(int)
	v2 := params[1].(int)
	return v1 % v2, nil
}

func Min(params []Value) (Value, error) {
	v1 := params[0].(int)
	v2 := params[1].(int)
	if v1 < v2 {
		return v1, nil
	}
	return v2, nil
}

func Max(params []Value) (Value, error) {
	v1 := params[0].(int)
	v2 := params[1].(int)
	if v1 < v2 {
		return v1, nil
	}
	return v2, nil
}

func Abs(params []Value) (Value, error) {
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
		if pt != wt {
			return nil, fmt.Errorf("wrong parameter type. got %v but expected %v", pt, wt)
		}
	}

	return sig.Func(params)
}
