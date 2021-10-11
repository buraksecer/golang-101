package main

import (
	"fmt"
	"time"
)

func methodF(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	methodF("A")

	go methodF("B")

	go func(message string) {
		fmt.Println(message)
	}("C")

	time.Sleep(time.Second)

	fmt.Println("FINISH")
}
