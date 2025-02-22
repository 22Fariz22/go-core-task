package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan uint8)
	defer close(ch1)

	numbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cancel := make(chan struct{})

	go func() {
		defer close(cancel)
		for _, v := range numbers {
			time.Sleep(100 * time.Millisecond)
			ch1 <- v
		}
	}()

	ch2 := convertToFloat64AndCube(cancel, ch1)

	out := collect(ch2)

	for v := range out {
		fmt.Println(v)
	}
}

func convertToFloat64AndCube(cancel chan struct{}, in <-chan uint8) chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for {
			select {
			case v := <-in:
				fl := uint8ToFloat64(v)
				sq := cube(fl)
				out <- sq
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func collect(in chan float64) chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for v := range in {
			out <- v
		}
	}()

	return out
}

func uint8ToFloat64(number uint8) float64 {
	return float64(number)
}
func cube(number float64) float64 {
	return number * number * number
}
