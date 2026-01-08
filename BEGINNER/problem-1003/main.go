// 1003 - Simple Sum

package main

import (
	"fmt"
)

func getNumbers() (int, int) {
	var A int
	var B int

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	return A, B
}
func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	a, b := getNumbers()
	x := add(a, b)
	fmt.Println("SOMA =", x)
}
