package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func main() {
	catalog := map[int]Product{
		101: {"Laptop", 1200.00},
		102: {"Mouse", 25.50},
		103: {"Keyboard", 80.00},
	}

	data, _ := json.MarshalIndent(catalog, "", "    ")
	fmt.Println(string(data))
}
