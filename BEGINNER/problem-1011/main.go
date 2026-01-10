package main

import "fmt"

func main() {
	var pi float64 = 3.14159
	var R int
	fmt.Scanf("%d\n", &R)
	var totalR = R * R * R

	var result float64 = (4.0 / 3.0) * pi * float64(totalR)

	fmt.Printf("VOLUME = %.3f\n", result)
}
