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
	if value != runtime.Int(4) {
		t.Errorf("Lambda 1 Failed got %v expected %v", value, 4)
	}
}

func Test_LambdaFromList(t *testing.T) {
	testStr := "var_a = [function var_b -> var_b * 2]; var_a.push(function var_b -> var_b * 3); \n var_a[0](2) + var_a[1](1)"

	value, err := execString(testStr, runtime.DefaultTimeout)
	if err != nil {
		t.Error(err)
	}
	if value != runtime.Int(7) {
		t.Errorf("Lambda 2 Failed got %v expected %v", value, 7)
	}
}

func Test_BuildInFuncs(t *testing.T) {
	tests := map[string]runtime.Value{
		"abs(0-3) == 3":    runtime.Bool(true),
		"mod(19, 5)":       runtime.Int(4),
		"min(max(1,3), 2)": runtime.Int(2),
		"pow(2,3)":         runtime.Int(8),
		"new_list(2)[1]":   runtime.Int(0),
		"random(1) == 0":   runtime.Bool(true),
		"time().is_list":   runtime.Bool(false),
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
		"1.is_list":                                                                               runtime.Bool(false),
		"[1].is_list":                                                                             runtime.Bool(true),
		"[1,2,3].length":                                                                          runtime.Int(3),
		"[].push(1) == [1]":                                                                       runtime.Bool(true),
		"[1].pop":                                                                                 runtime.Int(1),
		"[1].insert(0, 0) == [0, 1]":                                                              runtime.Bool(true),
		"[1].remove(0) == []":                                                                     runtime.Bool(true),
		"[3,2,1].sort == [1, 2, 3]":                                                               runtime.Bool(true),
		"[3,1,2].sort_with(function var_a, var_b -> var_a > var_b) == [3,2,1]":                    runtime.Bool(true),
		"[1].map(function var_a -> var_a * 2)[0]":                                                 runtime.Int(2),
		"var_a = [1]; var_b = var_a.copy; var_a.map(function var_a -> var_a * 2); var_a != var_b": runtime.Bool(true),
		"new_list(2).fill(2) == [2, 2]":                                                           runtime.Bool(true),
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
