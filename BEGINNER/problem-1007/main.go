// 1007 - Difference

package main

import (
	"fmt"
)

func getNumbers() (int, int, int, int) {
	var A, B, C, D int

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)
	fmt.Scanln(&D)

	return A, B, C, D
}
func diff(a int, b int, c int, d int) int {
	return (a*b - c*d)
}

func main() {
	a, b, c, d := getNumbers()
	x := diff(a, b, c, d)
	fmt.Printf("DIFERENCA = %d\n", x)
}
