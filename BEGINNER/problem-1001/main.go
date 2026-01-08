package main

import (
	"fmt"
)

func getNumbers() (int, int) {
	var a int
	var b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	return a, b
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	a, b := getNumbers()
	var x = add(a, b)
	fmt.Printf("X = %d\n", x)
}
