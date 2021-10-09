package main

import "fmt"

func main() {

	//Conditional type 1:
	if 10%3 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//Conditional type 2:
	if 10%2 == 0 {
		fmt.Println("true")
	}

	//Conditional type 3:
	if num := 31; num > 0 {
		fmt.Println(num, "> 0")
	} else if num > 32 {
		fmt.Println(num, "> 32")
	} else {
		fmt.Println(num, "else")
	}
}
