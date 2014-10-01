package main

import (
	"./runtime"
	"testing"
)

func Test_Equals(t *testing.T) {
	tests := map[string]runtime.Value{
		"var_a == var_b":                               true,
		"1 == 2":                                       false,
		"1 == [1]":                                     false,
		"[1, 1] == [1]":                                false,
		"var_a == [].length":                           true,
		"var_a = [[1,1], [1,1]]; var_a[0] == var_a[1]": true,
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
		"1 + 2 * 2": 5,
		"1 + 2 > 2": true,
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
		"var_a + var_b":             0,
		"var_a + 2":                 2,
		"var_a = [1]; 3 + var_a[0]": 4,
		"var_a - 10":                -10,
		"4 - 2 + 2":                 4,
		"1 + 1 + 1 - 3":             0,
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
