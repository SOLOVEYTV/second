package main

type MyInt = int // псевдоним

func main() {
	var a = 10
	var b MyInt = 20
	a = b // ок, так как это один и тот же тип
	b = a
}
