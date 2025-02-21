package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Cross(a, b))
}

func Cross(a, b []int) (bool, []int) {
	result := make([]int, 0, len(a)+len(b))

	mp := make(map[int]struct{})

	for _, v := range a {
		mp[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := mp[v]; ok {
			result = append(result, v)
		}
	}

	if len(result) > 0 {
		return true, result
	}
	return false, []int{}
}
