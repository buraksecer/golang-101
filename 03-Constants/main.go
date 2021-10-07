package main

import "fmt"

const FILENAME string = "/ROOT"
const NOTFOUND = 404
const (
	ERROR = 500
	SUCCESS = 200
)

func main() {
	fmt.Println(FILENAME)
	fmt.Println(NOTFOUND)
	fmt.Println(ERROR)
	fmt.Println(SUCCESS)
}
