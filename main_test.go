package main

import (
	"reflect"
	"testing"
)

func TestGetFibonacciArr(t *testing.T) {

	x := 2
	y := 10
	expected := []int{1, 2, 3, 5, 8, 13, 21, 34, 55}

	result := GetFibonacciArr(x, y)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect result. Expect %v, got %v", expected, result)
	}
}
