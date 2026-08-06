package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	person := Person{
		Name:    "Ваня",
		Age:     20,
		Address: "Спб, ул. Садовая",
	}

	fmt.Printf("%+v\n", person)

	data, _ := json.MarshalIndent(person, "", "  ")
	fmt.Println(string(data))

	//name := person.Name
	fmt.Println(person.Name)

	person.Name = "Дима"

	data, _ = json.MarshalIndent(person, "", "  ")
	fmt.Println(string(data))
}
