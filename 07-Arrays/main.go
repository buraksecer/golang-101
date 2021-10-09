package main

import (
	"fmt"
)

func main() {

	var intArray [5]int
	var strArray [5]string
	var boolArray [5]bool
	fmt.Println(intArray)
	fmt.Println(strArray)
	fmt.Println(boolArray)

	intArray[0] = 1
	strArray[0] = "0. index"
	boolArray[0] = true
	fmt.Println(intArray)
	fmt.Println(strArray)
	fmt.Println(boolArray)

	x := [5]int{0, 1, 2, 3, 4}
	fmt.Println(x)

	var y = [5]int{1, 2, 3}
	fmt.Println(y)

	k := [...]int{10, 20, 30, 44, 66, 11}
	fmt.Println(len(k))

	t := [5]int{1: 10, 3: 30}
	fmt.Println(t)

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  //data is passed by value
	strArray3 := &strArray1 //data is passed by refrence
	strArray1[0] = "Canada"
	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %v\n", *strArray3)

	var d2 [2][3]int
	fmt.Println(d2)

	var d3 [2][2][1]int
	fmt.Println(d3)

}
