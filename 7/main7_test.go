package main

import (
	"reflect"
	"testing"
)

func TestGenRange(t *testing.T) {
	tests := []struct {
		name    string
		start   int
		end     int
		lenWant int
		numbers []int
		close   bool
	}{
		{name: "4 positive numbers", start: 1, end: 5, lenWant: 4, numbers: []int{1, 2, 3, 4}, close: true},
		{name: "4 negative numbers ", start: -5, end: -1, lenWant: 4, numbers: []int{-5, -4, -3, -2}, close: true},
		{name: "0 numbers start:0 end:0", start: 0, end: 0, lenWant: 0, numbers: []int{}, close: true},
		{name: "0 numbers start:100 end:100", start: 100, end: 100, lenWant: 0, numbers: []int{}, close: true},
		{name: "1 numbers", start: 0, end: 1, lenWant: 1, numbers: []int{0}, close: true},
		{name: "wrong range", start: 5, end: 1, lenWant: 0, numbers: []int{}, close: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenRange(tt.start, tt.end)

			numbers := []int{}

			for v := range result {
				numbers = append(numbers, v)
			}

			lenNumbers := len(numbers)

			if lenNumbers != tt.lenWant {
				t.Errorf("Expected %d, got %d", tt.lenWant, lenNumbers)
			}

			_, ok := <-result
			if ok {
				t.Errorf("Expexted close channel %t, got %t", ok, tt.close)
			}

			if tt.start > tt.end && lenNumbers != tt.lenWant {
				t.Errorf("Expected %d, got %d", tt.lenWant, lenNumbers)
			}

			if tt.start > tt.end {
				_, ok := <-result
				if ok {
					t.Errorf("Expexted close channel %t, got %t", ok, tt.close)
				}

			}

			if !reflect.DeepEqual(tt.numbers, numbers) {
				t.Errorf("Expexted %v, got %v", tt.numbers, numbers)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name              string
		numberOfElements  int
		typeOfWantChannel reflect.Type
		close             bool
		numbers           []int
	}{
		{
			name:              "3 elements: 0,1,2",
			numberOfElements:  3,
			typeOfWantChannel: reflect.TypeOf(make(chan int)),
			close:             false,
			numbers:           []int{0, 1, 2},
		},
		{
			name:              "0 elements",
			numberOfElements:  0,
			typeOfWantChannel: reflect.TypeOf(make(chan int)),
			close:             false,
			numbers:           []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int, 10)
			for i := 0; i < tt.numberOfElements; i++ {
				ch <- i
			}
			close(ch)

			result := Merge(ch)

			numbers := []int{}
			for v := range result {
				numbers = append(numbers, v)
			}

			lenNumbers := len(numbers)
			if lenNumbers != tt.numberOfElements {
				t.Errorf("Expexted %d elements, got %d", tt.numberOfElements, lenNumbers)
			}

			_, ok := <-result
			if ok {
				t.Errorf("Expexted close channel %t, got %t", tt.close, ok)
			}

			typeOfResult := reflect.TypeOf(result)

			if typeOfResult != tt.typeOfWantChannel {
				t.Errorf("Expected type of channel %v, got %v", tt.typeOfWantChannel, typeOfResult)
			}

			if !reflect.DeepEqual(numbers, tt.numbers) {
				t.Errorf("Expexted numbers %v ,got %v", tt.numbers, numbers)
			}
		})
	}

}

/*
func Merge(channels ...<-chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for i := range channels {
		go func() {
			for val := range channels[i] {
				out <- val
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
*/
