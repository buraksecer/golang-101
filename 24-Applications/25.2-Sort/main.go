package main

import (
	"fmt"
	"sort"
)

var Series = []int{5, 4, 1, 3, 8, 0}
var Fruits = []string{"banana", "apple"}

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	fmt.Println(Series)
	sort.Ints(Series)
	fmt.Println(Series)

	fmt.Println(Fruits)
	sort.Strings(Fruits)
	fmt.Println(Fruits)

	s := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

	people := []Person{
		{"burak", 29},
		{"esad", 28},
		{"soner", 32},
		{"bünyamin", 31},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println(ByAge(people).Len())
	fmt.Println(ByAge(people).Less(0, 3))
	ByAge(people).Swap(0, 3)
	fmt.Println(people)
}
