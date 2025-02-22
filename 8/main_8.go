package main

import (
	"errors"
	"fmt"
	"sync/atomic"
)

func main() {
	var wg WaitGroupCastom

	fmt.Println(wg.counter)

	wg.Add()
	wg.Add()
	wg.Add()
	wg.Done()
	wg.Done()
	wg.Done()
	wg.counter--

	fmt.Println(wg.counter)
	wg.Wait()
	fmt.Println("wait is successfully done")
}

type WaitGroupCastom struct {
	counter uint64
}

func (w *WaitGroupCastom) Add() {
	// w.atom.Add(w.counter, 1)
	atomic.AddUint64(&w.counter, 1)
}

func (w *WaitGroupCastom) Done() error {
	if atomic.LoadUint64(&w.counter) == 0 {
		return errors.New("negative waitGroup counter")
	}

	atomic.SwapUint64(&w.counter, 1)
	return nil
}

func (w *WaitGroupCastom) Wait() {
	for atomic.LoadUint64(&w.counter) > 0 {
	}
}

/*
Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

* Напишите unit тесты к созданным функциям

*/
