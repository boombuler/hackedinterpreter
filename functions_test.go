package main

import (
	"./runtime"
	"testing"
)

func Test_LambdaSimple(t *testing.T) {
	testStr := "var_a = function var_b -> var_b * 2; \n var_a(2)"

	value, err := execString(testStr, runtime.DefaultTimeout)
	if err != nil {
		t.Error(err)
	}
	if value != 4 {
		t.Errorf("Lambda 1 failed got %v expected %v", value, 4)
	}
}

func Test_LambdaFromList(t *testing.T) {
	testStr := "var_a = [function var_b -> var_b * 2]; var_a.push(function var_b -> var_b * 3); \n var_a[0](2) + var_a[1](1)"

	value, err := execString(testStr, runtime.DefaultTimeout)
	if err != nil {
		t.Error(err)
	}
	if value != 7 {
		t.Errorf("Lambda 2 failed got %v expected %v", value, 7)
	}
}

func Test_Fibunacci(t *testing.T) {
	testStr := "function f1: var_a { if var_a < 3 { return 1 } f1(var_a-1) + f1(var_a-2) } f1(6)"

	value, err := execString(testStr, runtime.DefaultTimeout)
	if err != nil {
		t.Error(err)
	}
	if value != 8 {
		t.Errorf("Fibunacci Test failed got %v expected %v", value, 4)
	}
}

func Test_CustomFunctions(t *testing.T) {
	testStr := "function f1: var_a { var_a * 2 } [2].map(f1)[0]"

	value, err := execString(testStr, runtime.DefaultTimeout)
	if err != nil {
		t.Error(err)
	}
	if value != 4 {
		t.Errorf("Custom Functions Test failed got %v expected %v", value, 4)
	}
}

func Test_BuildInFuncs(t *testing.T) {
	tests := map[string]runtime.Value{
		"abs(0-3) == 3":    true,
		"mod(19, 5)":       4,
		"min(max(1,3), 2)": 2,
		"pow(2,3)":         8,
		"new_list(2)[1]":   0,
		"random(1) == 0":   true,
		"is_list(time())":  false,
		"is_list(1)":       false,
		"is_list([1])":     true,
	}
	for code, val := range tests {
		value, err := execString(code, runtime.DefaultTimeout)
		if err != nil {
			t.Error(err)
		}
		if !runtime.Equals(value, val) {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}

func Test_Methods(t *testing.T) {
	tests := map[string]runtime.Value{
		"[1,2,3].length":                                                                          3,
		"[].push(1) == [1]":                                                                       true,
		"[1].pop":                                                                                 1,
		"[1].insert(0, 0) == [0, 1]":                                                              true,
		"[1].remove(0) == []":                                                                     true,
		"[3,2,1].sort == [1, 2, 3]":                                                               true,
		"[3,1,2].sort_with(function var_a, var_b -> var_a > var_b) == [3,2,1]":                    true,
		"[1].map(function var_a -> var_a * 2)[0]":                                                 2,
		"var_a = [1]; var_b = var_a.copy; var_a.map(function var_a -> var_a * 2); var_a != var_b": true,
		"new_list(2).fill(2) == [2, 2]":                                                           true,
	}
	for code, val := range tests {
		value, err := execString(code, runtime.DefaultTimeout)
		if err != nil {
			t.Error(err)
		}
		if !runtime.Equals(value, val) {
			t.Errorf("Failed got %v expected %v\nTest:%v", value, val, code)
		}
	}
}
