package main

import (
	"sync/atomic"
)

func main() {
	wg := NewWG()
	wg.Add(1)
	wg.Done()
	wg.Wait()
}

type WaitGroupCastom struct {
	counter int64
}

func NewWG() *WaitGroupCastom {
	return &WaitGroupCastom{}

}

func (w *WaitGroupCastom) Add(incr int) {
	newVal := atomic.AddInt64(&w.counter, int64(incr))
	if newVal < 0 {
		panic("WaitGroup counter cannot be negative")
	}
}

func (w *WaitGroupCastom) Done() {
	newVal := atomic.AddInt64(&w.counter, int64(-1))
	if newVal < 0 {
		panic("WaitGroup counter cannot be negative")
	}
}

func (w *WaitGroupCastom) Wait() {
	for w.counter != 0 {
	}
}
