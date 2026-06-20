package main

import "fmt"

func main() {
	all := []string{"яблоко", "банан", "апельсин", "груша", "киви"}
	result := make([]string, 0, len(all)-1)
	for idx := range all {
		if idx == 2 {
			continue
		}

		result = append(result, all[idx])
	}
	fmt.Println(result)
}
