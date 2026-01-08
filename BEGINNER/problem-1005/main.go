// 1005 - Average 1

package main

import (
	"fmt"
)

func getNumbers() (float64, float64) {
	var A, B float64

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	return A, B
}
func avag(a float64, b float64) float64 {
	return (a*3.5 + b*7.5) / 11
}

func main() {
	a, b := getNumbers()
	x := avag(a, b)
	fmt.Printf("MEDIA = %.5f\n", x)
}
