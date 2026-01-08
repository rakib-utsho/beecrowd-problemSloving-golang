
// 1004 - Simple Product

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
func multi(a int, b int) int {
	multi := a * b
	return multi
}

func main() {
	a, b := getNumbers()
	x := multi(a, b)
	fmt.Println("PROD =", x)
}
