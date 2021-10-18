package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var jsonUser = []byte(`[{"Id":31,"Name":"Elma","Age":33}]`)

func main() {

	users := User{
		Id:   1,
		Name: "Burak",
		Age:  31,
	}

	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(b)

	var user []User
	errUser := json.Unmarshal(jsonUser, &user)
	if errUser != nil {
		fmt.Println("error:", errUser)
	}

	fmt.Println("\n", user)

	userList := []User{users, users}

	bs, errorList := json.Marshal(userList)
	if errorList != nil {
		fmt.Println(errorList)
	}

	fmt.Println(string(bs))
}
