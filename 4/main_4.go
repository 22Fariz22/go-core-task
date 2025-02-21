package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(Difference(slice1, slice2))
}

func Difference(slice1, slice2 []string) []string {
	result := make([]string, 0, len(slice1))

	mp := make(map[string]struct{})

	for _, v := range slice2 {
		mp[v] = struct{}{}
	}

	for _, v := range slice1 {
		if _, ok := mp[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

/*
На вход подаются два неупорядоченных слайса строк.
Например:
```
slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
slice2 := []string{"banana", "date", "fig"}
```
Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе,
но отсутствуют во втором.

* Напишите unit тесты к созданным функциям
*/
