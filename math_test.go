package main

import (
	"./runtime"
	"testing"
)

func Test_Equals(t *testing.T) {
	tests := map[string]runtime.Bool{
		"var_a == var_b":                               runtime.Bool(true),
		"1 == 2":                                       runtime.Bool(false),
		"1 == [1]":                                     runtime.Bool(false),
		"[1, 1] == [1]":                                runtime.Bool(false),
		"var_a == [].length":                           runtime.Bool(true),
		"var_a = [[1,1], [1,1]]; var_a[0] == var_a[1]": runtime.Bool(true),
	}
	for code, val := range tests {
		value, err := execString(code, runtime.DefaultTimeout)
		if err != nil {
			t.Error(err)
		}
		if value != val {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}

func Test_OperatorPrecedence(t *testing.T) {
	tests := map[string]runtime.Value{
		"1 + 2 * 2": runtime.Int(5),
		"1 + 2 > 2": runtime.Bool(true),
	}
	for code, val := range tests {
		value, err := execString(code, runtime.DefaultTimeout)
		if err != nil {
			t.Error(err)
		}
		if value != val {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}

func Test_AddSub(t *testing.T) {
	tests := map[string]runtime.Value{
		"var_a + var_b":             runtime.Int(0),
		"var_a + 2":                 runtime.Int(2),
		"var_a = [1]; 3 + var_a[0]": runtime.Int(4),
		"var_a - 10":                runtime.Int(-10),
		"4 - 2 + 2":                 runtime.Int(4),
		"1 + 1 + 1 - 3":             runtime.Int(0),
	}
	for code, val := range tests {
		value, err := execString(code, runtime.DefaultTimeout)
		if err != nil {
			t.Error(err)
		}
		if value != val {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}

func Test_AddSubFails(t *testing.T) {
	tests := []string{
		"[] + [1]", "var_a + true", "true + 3", "1 ++ 3",
		"false - true", "[1, 2, 3] + [3, 2, 1]",
	}
	for _, code := range tests {
		if _, err := execString(code, runtime.DefaultTimeout); err == nil {
			t.Errorf("Failed to fail! Test:%v", code)
		}
	}
}
