package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	visit([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println("call:", i)
	})

	xs := filter([]int{1, 2, 3, 4}, func(i int) bool {
		return i > 1
	})
	fmt.Println(xs)
}
