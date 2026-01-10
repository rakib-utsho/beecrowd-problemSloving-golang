// 1012 - Area
package main

import "fmt"

func main() {

	var A, B, C float64

	var pi float64 = 3.14159

	fmt.Scanf("%f %f %f\n", &A, &B, &C)

	rectangleTriangledArea := 0.5 * (A * C) // rectanle triangle area formula = 1/2 * base * height
	circleArea := pi * (C * C)              // circle area formula= pi*r^2
	trapeziumArea := 0.5 * (A + B) * C      // trapezium area formula = 1/2*(sum of parallel side)*height := 1/2*(a+b)*h
	squareArea := B * B                     // sqiuare area formula = s^2 := side*side
	rectangleArea := A * B                  // rectangle area formula = width*length := l*w

	fmt.Printf("TRIANGULO: %.3f\n", rectangleTriangledArea)
	fmt.Printf("CIRCULO: %.3f\n", circleArea)
	fmt.Printf("TRAPEZIO: %.3f\n", trapeziumArea)
	fmt.Printf("QUADRADO: %.3f\n", squareArea)
	fmt.Printf("RETANGULO: %.3f\n", rectangleArea)

}
