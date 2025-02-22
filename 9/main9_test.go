package main

import (
	"reflect"
	"testing"
)

func TestConvertToFloat64AndCube(t *testing.T) {
	tests := []struct {
		name             string
		quantityElements int
		numbersIn        []uint8
		numbersOut       []float64
		closeCancelChan  bool
		closeOutChan     bool
	}{
		{
			name:             "success 5 element",
			quantityElements: 5,
			numbersIn:        []uint8{1, 2, 3, 4, 5},
			numbersOut:       []float64{1, 8, 27, 64, 125},
			closeOutChan:     false,
			closeCancelChan:  false,
		},
		{
			name:             "success 0 element",
			quantityElements: 0,
			numbersIn:        []uint8{},
			numbersOut:       []float64{},
			closeOutChan:     false,
			closeCancelChan:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch1 := make(chan uint8)
			defer close(ch1)

			cancel := make(chan struct{})

			go func() {
				defer close(cancel)
				for _, v := range tt.numbersIn {
					ch1 <- v
				}
			}()

			ch2 := convertToFloat64AndCube(cancel, ch1)

			result := make([]float64, 0, len(tt.numbersIn))

			for v := range ch2 {
				result = append(result, v)
			}

			lenResult := len(result)

			if lenResult != int(tt.quantityElements) {
				t.Errorf("Expected quantity elements %d, got %d", lenResult, tt.quantityElements)
			}

			_, ok := <-cancel
			if ok {
				t.Errorf("Expected close cancel channel %t, got %t", tt.closeCancelChan, ok)
			}

			_, ok = <-ch2
			if ok {
				t.Errorf("Expected close ch2 %t, got %t", tt.closeOutChan, ok)
			}

		})
	}
}

func TestCollect(t *testing.T) {
	tests := []struct {
		name             string
		quantityElements int
		numbersIn        []float64
		numbersOut       []float64
		closeOutChan     bool
	}{
		{
			name:             "success 5 element",
			quantityElements: 5,
			numbersIn:        []float64{1, 2, 3, 4, 5},
			numbersOut:       []float64{1, 8, 27, 64, 125},
			closeOutChan:     false,
		},
		{
			name:             "success 0 element",
			quantityElements: 0,
			numbersIn:        []float64{},
			numbersOut:       []float64{},
			closeOutChan:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch1 := make(chan float64)

			go func() {
				defer close(ch1)
				for _, v := range tt.numbersIn {
					ch1 <- v
				}
			}()

			ch2 := collect(ch1)

			result := make([]float64, 0, 5)

			for v := range ch2 {
				result = append(result, v)
			}

			lenResult := len(result)

			if lenResult != int(tt.quantityElements) {
				t.Errorf("Expected quantity elements %d, got %d", lenResult, tt.quantityElements)
			}

			_, ok := <-ch2
			if ok {
				t.Errorf("Expected close ch2 %t, got %t", tt.closeOutChan, ok)
			}

		})
	}

}

func TestUint8ToFloat64(t *testing.T) {
	tests := []struct {
		name      string
		inNumber  uint8
		outNumber float64
	}{
		{name: "success number 0", inNumber: 0, outNumber: 0},
		{name: "success number 5", inNumber: 5, outNumber: 5},
		{name: "success number 1000", inNumber: 100, outNumber: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uint8ToFloat64(tt.inNumber)

			typeOut := reflect.TypeOf(tt.outNumber)
			typeResult := reflect.TypeOf(result)

			if typeOut != typeResult {
				t.Errorf("Expected type %v, got %v", typeOut, typeResult)
			}
		})
	}
}

func TestCube(t *testing.T) {
	tests := []struct {
		name string
		in   float64
		out  float64
	}{
		{name: "success 0", in: 0, out: 0},
		{name: "success 1", in: 1, out: 1},
		{name: "success 1", in: 2, out: 8},
		{name: "success 1", in: 5, out: 125},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := cube(tt.in)
			typeOfIn := reflect.TypeOf(tt.in)
			typeOfOut := reflect.TypeOf(out)

			if typeOfIn != typeOfOut {
				t.Errorf("Expected type %v, got %v", typeOfIn, typeOfOut)
			}
		})
	}
}
