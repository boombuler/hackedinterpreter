package main

import (
	"errors"
	"github.com/boombuler/gold"
	"strconv"
)

type Value interface{}

func ConstIntValue(t *gold.Token, c *Context) (Value, error) {
	if t.Rule == Rule_ValueNumber {
		return ConstIntValue(t.Tokens[0], c)
	}
	val, err := strconv.Atoi(t.Text)
	return val, err
}

func ConstBoolValue(t *gold.Token, c *Context) (Value, error) {
	if t.Rule == Rule_Value2 {
		return ConstBoolValue(t.Tokens[0], c)
	}
	if t.Rule == Rule_BoolTrue {
		return true, nil
	}
	if t.Rule == Rule_BoolFalse {
		return false, nil
	}
	return nil, errors.New("unknown bool constant")
}

func getTypeName(v Value) string {
	switch v.(type) {
	case int:
		return "Integer"
	case bool:
		return "Boolean"
	case *List:
		return "List"
	case *Function:
		return "Function"
	default:
		return "undefined"
	}
}
