// 1008 - Salary

package main

import (
	"fmt"
)

func getNumbers() (int, int) {
	var A, B int

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	return A, B
}
func getSalary() float64 {
	var salary float64
	fmt.Scanln(&salary)

	return salary
}
func diff(b int, c float64) float64 {
	return float64(b) * c
}

func main() {
	a, b := getNumbers()
	c := getSalary()
	x := diff(b, c)

	fmt.Printf("NUMBER = %d\n", a)
	fmt.Printf("SALARY = U$ %.2f\n", x)
}
