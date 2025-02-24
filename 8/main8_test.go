package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewWG(t *testing.T) {
	result := NewWG()

	typeOfResult := reflect.TypeOf(result)
	wantTypeOfWG := reflect.TypeOf(&WaitGroupCastom{})

	if typeOfResult != wantTypeOfWG {
		t.Errorf("Expected type %v,got %v", wantTypeOfWG, typeOfResult)
	}

}

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		add         int
		wantCounter int
	}{
		{name: "add 1", add: 1, wantCounter: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &WaitGroupCastom{}
			wg.Add(tt.add)

			result := wg.counter
			if result != wg.counter {
				t.Errorf("Expected counter %d, got %d", tt.wantCounter, result)
			}

		})
	}
}

func TestAddWithPanic(t *testing.T) {
	wg := NewWG()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Exeption panic, got %v", r)
		}
	}()

	wg.Add(-1)
}

func TestDone(t *testing.T) {
	tests := []struct {
		name        string
		add         int
		done        int
		wantCounter int
	}{
		{name: "add 1, done 1", add: 1, done: 1, wantCounter: 0},
		{name: "add 3 done 2", add: 3, done: 2, wantCounter: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &WaitGroupCastom{}
			wg.Add(tt.add)

			for i := 0; i < tt.done; i++ {
				wg.Done()
			}

			result := wg.counter
			if result != wg.counter {
				t.Errorf("Expected counter %d, got %d", tt.wantCounter, result)
			}
		})
	}
}

func TestDoneWithPanic(t *testing.T) {
	wg := NewWG()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Exeption panic, got %v", r)
		}
	}()

	wg.Add(0)
	wg.Done()
}

func TestWait(t *testing.T) {
	wg := NewWG()

	start := time.Now()
	wg.Wait()
	duration := time.Since(start)

	if duration > 10*time.Millisecond {
		t.Errorf("Wait() took too long time %v when counter is 0", duration)
	}

}
