package main

import (
	"fmt"
	"maps"
)

func main() {
	m := NewStringIntMap()

	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)
	fmt.Println(m.mp)

	m.Remove("two")
	fmt.Println(m.mp)

	fmt.Println(m.Copy())

	fmt.Println(m.Exists("five"))
	fmt.Println(m.Exists("one"))

	fmt.Println(m.Get("five"))
	fmt.Println(m.Get("one"))
}

type StringIntMap struct {
	mp map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{mp: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.mp[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.mp, key)
}

func (m *StringIntMap) Copy() map[string]int {
	newMp := make(map[string]int)
	maps.Copy(newMp, m.mp)
	return newMp
}

func (m *StringIntMap) Exists(key string) bool {
	if _, ok := m.mp[key]; !ok {
		return false
	}
	return true
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.mp[key]
	return v, ok
}
