package main

import (
	"bytes"
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

type ValueType string

const (
	INT       ValueType = "Integer"
	BOOL      ValueType = "Boolean"
	LIST      ValueType = "List"
	FUNC      ValueType = "Function"
	ANY       ValueType = "Anything"
	DRAWINGSF ValueType = "drawing surface"
)

func getTypeName(v Value) ValueType {
	switch v.(type) {
	case int:
		return INT
	case bool:
		return BOOL
	case *List:
		return LIST
	case *Function:
		return FUNC
	default:
		return ANY
	}
}

func ToString(v Value) string {
	var ToStringNested func(v Value, parentLists []*List) string
	ToStringNested = func(v Value, parentLists []*List) string {
		switch v.(type) {
		case ValueType:
			return string(v.(ValueType)) // drawing surface etc.
		case int:
			return strconv.Itoa(v.(int))
		case bool:
			if v.(bool) {
				return "true"
			} else {
				return "false"
			}
		case *List:
			{
				// check if list is in parentList
				l := v.(*List)
				for i, pl := range parentLists {
					if l == pl {
						if i < len(parentLists)-1 {
							return "(parent Collection)"
						} else {
							return "(this Collection)"
						}
					}
				}
				curPL := make([]*List, len(parentLists), len(parentLists)+1)
				copy(curPL, parentLists)
				curPL = append(curPL, l)

				buf := &bytes.Buffer{}
				buf.WriteRune('[')

				for i, val := range l.content {
					if i != 0 {
						buf.WriteString(", ")
					}

					buf.WriteString(ToStringNested(val, curPL))
				}

				buf.WriteRune(']')
				return buf.String()
			}
		case *Function:
			return string(FUNC)
		default:
			return ""
		}
	}

	return ToStringNested(v, []*List{})
}
