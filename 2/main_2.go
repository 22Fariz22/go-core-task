package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := NewRandomSlice(100)
	fmt.Println(arr)

	odd := SliceExample(arr)
	fmt.Println(odd)

	addedNumber := AddElements(arr, 777)
	fmt.Println(addedNumber)

	copyArr := CopySlice(arr)
	fmt.Println(copyArr)

	removed := RemoveElement(arr, 2)
	fmt.Println(removed)
}

func NewRandomSlice(rangeInt int) []int {
	arr := []int{}

	for len(arr) < 10 {
		random := rand.Intn(rangeInt)
		arr = append(arr, random)
	}
	return arr
}

func SliceExample(arr []int) []int {
	oddArr := []int{}

	for _, v := range arr {
		if v%2 == 0 {
			oddArr = append(oddArr, v)
		}
	}

	return oddArr
}

func AddElements(arr []int, number int) []int {
	newArr := make([]int, len(arr))

	copy(newArr, arr)
	newArr = append(newArr, number)

	return newArr
}

func CopySlice(originalArr []int) []int {
	copyArr := make([]int, len(originalArr))
	copy(copyArr, originalArr)
	return copyArr
}

func RemoveElement(originalSlice []int, index int) []int {
	if index >= len(originalSlice) || index < 0 {
		return originalSlice
	}

	newSlice := make([]int, len(originalSlice)-1)

	copy(newSlice[:index], originalSlice[:index])
	copy(newSlice[index:], originalSlice[index+1:])

	return newSlice
}
