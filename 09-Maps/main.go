package main

import "fmt"

func main() {

	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee)

	employee = make(map[string]int)
	fmt.Println(employee)

	employee["Mark"] = 31 // Update element
	employee["Sandy"] = 40
	fmt.Println(employee)
	fmt.Println(employee["Mark"])

	employee["Rocky"] = 30 // Add element
	fmt.Println(employee)

	delete(employee, "Mark") // Delete element
	fmt.Println(employee)
}
