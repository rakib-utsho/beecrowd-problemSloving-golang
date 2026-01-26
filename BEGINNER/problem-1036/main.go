package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, t float64
	fmt.Scanf("%f %f %f\n", &a, &b, &c)

	if (b*b)-(4*a*c) < 0 || a == 0 {
		fmt.Println("Impossivel calcular")
	} else {
		t = math.Sqrt((b * b) - (4 * a * c))
		R1 := ((-b + t) / (2 * a))
		R2 := ((-b - t) / (2 * a))

		fmt.Printf("R1 = %.5f\n", R1)
		fmt.Printf("R2 = %.5f\n", R2)
	}
}
