package main

import "fmt"

func main() {

	defer second()
	first()
	first()
	first()
	first()
	first()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
