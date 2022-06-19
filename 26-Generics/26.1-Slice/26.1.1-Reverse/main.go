package main

import "fmt"

func main() {
	var list = []int{1, 2, 3, 4}
	Reverse(list)
	fmt.Println(list)
}

func Reverse[T any](list []T) {
	first := 0
	var last = len(list) - 1
	for first < last {
		list[first], list[last] = list[last], list[first]
		first++
		last--
	}
}
