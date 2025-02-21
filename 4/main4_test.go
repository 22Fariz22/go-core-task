package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := Difference(slice1, slice2)

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
