package main

import (
	"maps"
	"reflect"
	"testing"
)

func TestNewStringMap(t *testing.T) {
	newStruct := NewStringIntMap()

	expectedStruct := &StringIntMap{mp: make(map[string]int)}

	reflectNewStruct := reflect.TypeOf(newStruct)
	reflectExpected := reflect.TypeOf(expectedStruct)

	if reflectNewStruct != reflectExpected {
		t.Errorf("expected type %v, got %v", reflectExpected, reflectNewStruct)
	}
}

func TestAdd(t *testing.T) {
	newStruct := NewStringIntMap()
	key := "one"
	value := 1
	newStruct.Add(key, value)

	v, ok := newStruct.mp[key]
	if v != value || !ok {
		t.Errorf("Expected %d or %t, got %d or %t", value, true, v, ok)
	}

}
func TestRemove(t *testing.T) {
	newStruct := NewStringIntMap()
	key := "one"
	value := 1
	newStruct.mp[key] = value
	newStruct.Remove(key)

	_, ok := newStruct.mp[key]
	if ok {
		t.Errorf("Expected %t false, got %t", false, ok)
	}
}

func TestCopy(t *testing.T) {
	stringIntMp := &StringIntMap{mp: make(map[string]int)}

	stringIntMp.mp["one"] = 1
	stringIntMp.mp["two"] = 2
	stringIntMp.mp["three"] = 3

	copyMp := stringIntMp.Copy()

	if !maps.Equal(stringIntMp.mp, copyMp) {
		t.Errorf("Expected %v, got %v", stringIntMp.mp, copyMp)
	}
}

func TestExists(t *testing.T) {
	stringIntMp := &StringIntMap{mp: make(map[string]int)}
	key := "one"
	value := 1

	stringIntMp.mp[key] = value

	exist := stringIntMp.Exists(key)

	if !exist {
		t.Errorf("Expected %t, got %t", true, exist)
	}

	key = "ten"
	exist = stringIntMp.Exists(key)
	if exist {
		t.Errorf("Expected %t, got %t", false, exist)
	}

}

func TestGet(t *testing.T) {
	stringIntMp := &StringIntMap{mp: make(map[string]int)}
	key := "one"
	value := 1

	stringIntMp.mp[key] = value
	v, ok := stringIntMp.Get(key)
	if v != value {
		t.Errorf("Expected %d , got %d ", value, v)
	}
	if !ok {
		t.Errorf("Expected %t, got %t", true, ok)
	}

}
