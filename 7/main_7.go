package main

import (
	"fmt"
	"sync"
)

func main() {
	in1 := GenRange(0, 3)
	in2 := GenRange(3, 5)
	in3 := GenRange(9, 7)
	out := Merge(in1, in2, in3)

	for v := range out {
		fmt.Println(v)
	}
}

func GenRange(start, end int) <-chan int {
	out := make(chan int)

	if start > end {
		close(out)
		return out
	}

	go func() {
		defer close(out)

		for i := start; i < end; i++ {
			out <- i
		}
	}()
	return out
}

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
