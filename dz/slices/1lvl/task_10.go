package main

import "fmt"

func main() {
	all := []string{"Go", "Java", "Python", "Ruby"}

	findAndPrintOne("Python", all)
}

func findAndPrintOne(findValue string, list []string) {
	for _, value := range list {
		if value == findValue {
			fmt.Printf("%s - Найден\n", findValue)

			return
		}
	}

	fmt.Printf("%s - Не найден\n", findValue)
}
