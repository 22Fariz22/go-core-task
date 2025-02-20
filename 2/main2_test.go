package main

import (
	"reflect"
	"testing"
)

func TestNewRandomSlice(t *testing.T) {
	rangeRandom := 100
	result := NewRandomSlice(rangeRandom)

	if len(result) != 10 {
		t.Errorf("Expected len slice=%d, got %d", 10, len(result))
	}
}

func TestSliceExample(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := SliceExample(arr)

	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := AddElements(arr, 6)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveElement(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	t.Run("succes removed", func(t *testing.T) {
		index := 3
		result := RemoveElement(arr, index)

		expepted := []int{1, 2, 3, 5}

		if !reflect.DeepEqual(result, expepted) {
			t.Errorf("Expected %v, got %v", expepted, result)
		}
	})

	t.Run("index is greater than slice size", func(t *testing.T) {
		index := 7
		result := RemoveElement(arr, index)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("index is less than slice size", func(t *testing.T) {
		index := -2
		result := RemoveElement(arr, index)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestCopySlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := CopySlice(arr)

	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
