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

func ExpressionEqual(t *gold.Token, c *Context) (Value, error) {
	v1, err := Exec(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	v2, err := Exec(t.Tokens[2], c)
	if err != nil {

	}
	if getTypeName(v1) == getTypeName(v2) {
		return v1 == v2, nil
	}
	return false, nil
}

func ExpressionNotEqual(t *gold.Token, c *Context) (Value, error) {
	v1, err := Exec(t.Tokens[0], c)
	if err != nil {
		return nil, err
	}
	v2, err := Exec(t.Tokens[2], c)
	if err != nil {

	}
	if getTypeName(v1) == getTypeName(v2) {
		return v1 != v2, nil
	}
	return true, nil
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
