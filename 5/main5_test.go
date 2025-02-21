package main

import (
	"reflect"
	"testing"
)

func TestCross(t *testing.T) {
	tests := []struct {
		name     string
		sliceA   []int
		sliceB   []int
		expected []int
		exist    bool
		lenght   int
	}{
		{
			name:     "have common elements",
			sliceA:   []int{65, 3, 58, 678, 64},
			sliceB:   []int{64, 2, 3, 43},
			expected: []int{64, 3},
			exist:    true,
			lenght:   2,
		},
		{
			name:     "haven`t common elements",
			sliceA:   []int{65, 58, 678},
			sliceB:   []int{64, 2, 3, 43},
			expected: []int{},
			exist:    false,
			lenght:   0,
		},
		{
			name:     "haven`t common elements, because both slices haven`t elements",
			sliceA:   []int{},
			sliceB:   []int{},
			expected: []int{},
			exist:    false,
			lenght:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultExist, result := Cross(tt.sliceA, tt.sliceB)
			lenghtResult := len(result)

			if !reflect.DeepEqual(result, tt.expected) && resultExist != tt.exist {
				t.Errorf("expected exist:%t  result:%v lenght:%d, got  exist:%t  result:%v lenght:%d",
					tt.exist, tt.expected, tt.lenght, resultExist, result, lenghtResult)
			}

		})
	}
}
