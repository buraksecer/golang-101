package main

import "fmt"

func Connection() {
	fmt.Println("Success!")
}

func Close() {
	fmt.Println("Close connection!")
}

func Add() {
	defer Close()
	Connection()
}

func main() {
	Add()
}
