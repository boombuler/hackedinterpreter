package main

import (
	"github.com/boombuler/gold"
)

func AddValues(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantInt(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 + vi2), nil
}

func SubValues(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantInt(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 - vi2), nil
}

func MultiplyValues(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantInt(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 * vi2), nil
}

func DivideValues(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantInt(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 / vi2), nil
}

func ParenValue(t *gold.Token, c *Context) (Value, error) {
	return Exec(t.Tokens[1], c)
}

func ExpressionGt(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantInt(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 > vi2), nil
}

func ExpressionLt(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantInt(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantInt(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 < vi2), nil
}

func compare(v1, v2 Value) bool {
	v1L, ok1 := v1.(*List)
	v2L, ok2 := v2.(*List)
	if ok1 != ok2 {
		return false
	}
	if ok1 {
		if len(v1L.content) == len(v2L.content) {
			for i := 0; i < len(v1L.content); i++ {
				if !compare(v1L.content[i], v2L.content[i]) {
					return false
				}
			}
			return true
		}
		return false
	}
	if getTypeName(v1) == getTypeName(v2) {
		return v1 == v2
	}
	return false
}

func ExpressionEqual(t *gold.Token, c *Context) (Value, error) {
	v1, err := Exec(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	v2, err := Exec(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return compare(v1, v2), nil
}

func ExpressionNotEqual(t *gold.Token, c *Context) (Value, error) {
	v1, err := Exec(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	v2, err := Exec(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return !compare(v1, v2), nil
}

func ExpressionLogicalAND(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantBool(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantBool(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 && vi2), nil
}

func ExpressionLogicalOR(t *gold.Token, c *Context) (Value, error) {
	vi1, err := ExecWantBool(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	vi2, err := ExecWantBool(t.Tokens[2], c)
	if err != nil {
		return nil, err
	}

	return Value(vi1 || vi2), nil
}
