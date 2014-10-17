package runtime

import (
	"bytes"
	"fmt"
)

type Value interface{}
type ValueType string

const (
	INT       ValueType = "Integer"
	BOOL      ValueType = "Boolean"
	LIST      ValueType = "List"
	FUNC      ValueType = "Function"
	ANY       ValueType = "Anything"
	DRAWINGSF ValueType = "drawing surface"
)

func GetType(v Value) ValueType {
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

func callList(c Callable, ctx *Context) (*List, error) {
	val, err := ctx.Call(c)
	if err != nil {
		return nil, err
	}
	list, ok := val.(*List)
	if !ok {
		return nil, fmt.Errorf("List expected got %v", val)
	}
	return list, nil
}

func callInt(c Callable, ctx *Context) (int, error) {
	val, err := ctx.Call(c)
	if err != nil {
		return 0, err
	}
	iv, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("Integer expected got %v", val)
	}
	return iv, nil
}

func callBool(c Callable, ctx *Context) (bool, error) {
	val, err := ctx.Call(c)
	if err != nil {
		return false, err
	}
	bv, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("Boolean expected got %v", val)
	}
	return bv, nil
}

func NewAddToValues(valSlice, value Callable) Callable {
	return newCallable(nil, func(c *Context) (Value, error) {
		sliceV, err := c.Call(valSlice)
		if err != nil {
			return nil, err
		}
		slice := sliceV.([]Value)

		val, err := c.Call(value)
		if err != nil {
			return nil, err
		}

		return append(slice, val), nil
	}, valSlice, value)
}

func NewValues(value Callable) Callable {
	return newCallable(nil, func(c *Context) (Value, error) {
		val, err := c.Call(value)
		if err != nil {
			return nil, err
		}

		return []Value{val}, nil
	}, value)
}

func Equals(v1, v2 Value) bool {
	v1L, ok1 := v1.(*List)
	v2L, ok2 := v2.(*List)
	if ok1 != ok2 {
		return false
	}
	if ok1 {
		if len(v1L.content) == len(v2L.content) {
			for i := 0; i < len(v1L.content); i++ {
				if !Equals(v1L.content[i], v2L.content[i]) {
					return false
				}
			}
			return true
		}
		return false
	}
	if GetType(v1) == GetType(v2) {
		return v1 == v2
	}
	return false
}

func ToString(v Value) string {
	var ToStringNested func(v Value, parentLists []*List) string
	ToStringNested = func(v Value, parentLists []*List) string {
		switch v.(type) {
		case ValueType:
			return string(v.(ValueType)) // drawing surface etc.
		case int:
			return fmt.Sprintf("%v", v.(int))
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
