package main

import "fmt"

func main() {

	//Loops type 1:
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//Loops type 2:
	x := 4
	for ; x <= 5; x++ {
		fmt.Println(x)
	}

	//Loops type 3:
	y := 5
	for y <= 7 {
		fmt.Println(y)
		y++
	}

	//Loops type 4:
	for z := 7; ; z++ {
		fmt.Println(z)
		if z == 10 {
			break
		}
	}

	//Loops type 5:
	strName := map[int]string{1: "Burak", 2: "Son", 3: "Esad"}
	for index, element := range strName {
		fmt.Println("Index :", index, " Element :", element)
	}

	//Loops type 6:
	for key := range strName {
		fmt.Println(key)
	}

	//Loops type 7: (like foreach)
	for _, value := range strName {
		fmt.Println(value)
	}

	//Loops type 8:
	for range "12345" {
		fmt.Println("*")
	}

	//Loops type 9: (slice)
	s := []string{"Burak", "Banu"}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
