package main

import (
	"fmt"
)

type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() {
	for _, v := range g {
		fmt.Println(v)
	}
}

func main() {
	g := GenericSlice[string]{"burak", "secer", "washere"}
	gg := GenericSlice[int]{31, 40, 34}
	g.Print()
	gg.Print()
}
