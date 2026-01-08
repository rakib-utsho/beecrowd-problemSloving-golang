package main

import (
	"fmt"
)

func getNumber() float64 {
	var n float64
	fmt.Scanf("%g", &n)
	return n
}

func area(n float64) float64 {
	var pai float64 = 3.14159
	var area float64
	var R float64 = n * n
	area = pai * R

	return area
}
func main() {
	var n = getNumber()
	var a = area(n)
	fmt.Printf("A=%.4f\n", a)
}
