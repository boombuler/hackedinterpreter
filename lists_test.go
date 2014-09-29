package main

import (
	"./runtime"
	"testing"
)

func Test_Lists(t *testing.T) {
	tests := map[string]runtime.Value{
		"new_list(1).map(function var_a -> var_a + 2)[0]": runtime.Int(2),
		"[0, 1, 2].length":                                runtime.Int(3),
		"[0, 1, 2][1]":                                    runtime.Int(1),
		"var_a = [1];\nvar_a[0] = 2;\nvar_a[0]":           runtime.Int(2),
		"[[1, 2, 3]].push(1).length":                      runtime.Int(2),
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
