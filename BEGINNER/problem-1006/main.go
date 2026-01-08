// 1006 - Average 2

package main

import (
	"fmt"
)

func getNumbers() (float64, float64, float64) {
	var A, B, C float64

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	return A, B, C
}
func avag(a float64, b float64, c float64) float64 {
	return (a*2 + b*3 + c*5) / 10
}

func main() {
	a, b, c := getNumbers()
	x := avag(a, b, c)
	fmt.Printf("MEDIA = %.1f\n", x)
}
