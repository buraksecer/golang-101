package main

import (
	"fmt"
	"reflect"
)

func main() {

	//To declare the type for a variable that holds a slice, use an empty pair of square brackets, followed by the type of elements the slice will hold.
	var intSlice []int
	var strSlice []string
	fmt.Println(intSlice)
	fmt.Println(strSlice)

	//Slice can be created using the built-in function make. When you use make, one option you have is to specify the length of the slice. When you just specify the length, the capacity of the slice is the same.
	intSlice = make([]int, 10)//when length and capacity is same
	strSlice = make([]string, 10, 20)//when length and capacity is different
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	//A slice literal contain empty brackets followed by the type of elements the slice will hold, and a list of the initial values each element will have in curly braces.
	intSlice = []int{10, 20, 30, 40}
	strSlice = []string{"India", "Canada", "Japan"}
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	//A slice can be declare using new keyword followed by capacity in square brackets then type of elements the slice will hold.
	intSlice = new([50]int)[0:10]
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	//To add an item to the end of the slice, use the append() method.
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
	a = append(a, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	//You access the slice items by referring to the index number.
	intSlice = []int{10, 20, 30, 40}
	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[1:4])

	strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 3)
	fmt.Println(strSlice)

	o := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(o), cap(a))

	b := make([]int, 5, 10) // Create a bigger slice
	copy(b, o) // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(o), cap(b))

	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}
	slice2 = append(slice2, slice1...)
	fmt.Println(slice2)

	strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strSlice, "Canada"))
	fmt.Println(itemExists(strSlice, "Africa"))
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
