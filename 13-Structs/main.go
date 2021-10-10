package main

import "fmt"

/*type identifier struct{
	field1 data_type
	field2 data_type
	field3 data_type
}*/

type Employee struct {
	Name string
	Age  int
}

func (obj *Employee) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}

func main() {

	type sports struct {
		capacity float64
		name     string
		members  struct {
			id       int
			fullName string
		}
	}
	sport := sports{
		capacity: 2,
		name:     "tennis",
		members: struct {
			id       int
			fullName string
		}{1, "burak"},
	}
	fmt.Println(sport)

	type colors struct {
		name string
	}
	type car struct {
		id    int
		color []colors
	}
	newCar := car{
		id: 1,
		color: []colors{
			{name: "red"},
			{name: "black"},
		},
	}
	fmt.Println(newCar)

	emp1 := Employee{Name: "Mr. Fred"}
	emp1.Info()
	fmt.Println(emp1)
}


