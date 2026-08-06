package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Grade int
}

func main() {
	students := []Student{
		{
			Name:  "Alice",
			Grade: 85,
		},
		{
			Name:  "Bob",
			Grade: 85,
		},
		{
			Name:  "Bob",
			Grade: 85,
		},
	}

	data, _ := json.MarshalIndent(students, "", "    ")
	fmt.Println(string(data))
}
