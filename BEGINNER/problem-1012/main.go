package main

import "fmt"

func main() {

	var A, B, C float64

	var pi float64 = 3.14159

	fmt.Scanf("%f %f %f\n", &A, &B, &C)

	rectangleTriangledArea := 0.5 * (A * C)
	circleArea := pi * (C * C)
	trapeziumArea := 0.5 * (A + B) * C
	squareArea := B * B
	rectangleArea := A * B

	fmt.Printf("TRIANGULO: %.3f\n", rectangleTriangledArea)
	fmt.Printf("CIRCULO: %.3f\n", circleArea)
	fmt.Printf("TRAPEZIO: %.3f\n", trapeziumArea)
	fmt.Printf("QUADRADO: %.3f\n", squareArea)
	fmt.Printf("RETANGULO: %.3f\n", rectangleArea)

}
