package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "Hello World"
	freq := CharFrequency(text)
	fmt.Println(freq) // map[h:1 e:1 l:3 o:2 w:1 r:1 d:1]
	ch, cnt := MostCommonChar(text)
	fmt.Printf("%c: %d\n", ch, cnt) // l: 3
}

func CharFrequency(s string) map[rune]int {
	temp := make(map[rune]int)
	for _, char := range s {
		char = unicode.ToLower(char)
		if unicode.IsSpace(char) {
			continue
		}
		temp[char]++
	}

	return temp
}

func MostCommonChar(s string) (rune, int) {
	temp := CharFrequency(s)
	var maxCount int
	var maxChar rune
	for char, count := range temp {
		if count > maxCount || (count == maxCount && char < maxChar) {
			maxCount = count
			maxChar = char
		}
	}

	return maxChar, maxCount
}
