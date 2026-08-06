package main

import (
	"encoding/json"
	"fmt"
)

type User2 struct {
	Name    string
	Age     int
	Address string
}

func main() {
	user := &User2{
		Name:    "Дима",
		Age:     20,
		Address: "Спб",
	}

	updateAge(user, 30)

	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))

	user.updateUserAge(30)
}

func updateAge(user *User2, newAge int) {
	user.Age = newAge
}

func (user *User2) updateUserAge(newAge int) {
	user.Age = newAge
}
