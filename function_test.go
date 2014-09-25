package main

import (
	"testing"
)

func Test_LambdaSimple(t *testing.T) {
	testStr := "var_a = function var_b -> var_b * 2; \n var_a(2)"

	value, err := ExecuteString(testStr)
	if err != nil {
		t.Error(err)
	}
	if value != 4 {
		t.Errorf("Lambda 1 Failed got %v expected %v", value, 4)
	}
}

func Test_LambdaFromList(t *testing.T) {
	testStr := "var_a = [function var_b -> var_b * 2]; var_a.push(function var_b -> var_b * 3); \n var_a[0](2) + var_a[1](1)"

	value, err := ExecuteString(testStr)
	if err != nil {
		t.Error(err)
	}
	if value != 7 {
		t.Errorf("Lambda 1 Failed got %v expected %v", value, 7)
	}
}
