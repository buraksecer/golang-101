package main

import "fmt"

func standard(param int) {
	param = 31
}

func withPointer(param *int) {
	*param = 31
}

func main() {

	i := 1
	fmt.Println("First value:", i)
	standard(i)
	fmt.Println("Standard value:", i)
	withPointer(&i)
	fmt.Println("With pointer value:", i)
}
