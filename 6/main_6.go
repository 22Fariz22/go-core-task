package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})
	n := 10
	max := 100

	numbers := RandomGenerator(n, max)

	go func() {
		for v := range numbers {
			fmt.Println(v)
		}
		done <- struct{}{}

	}()
	<-done
}

func RandomGenerator(n, max int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < n; i++ {
			out <- r.Intn(100)
		}
	}()

	return out
}
