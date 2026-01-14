// beecrowd | 1014 Consumption

package main

import "fmt"

func main() {
	var X int
	var Y float64

	fmt.Scanf("%d\n", &X)
	fmt.Scanf("%f\n", &Y)

	var consumption float64 = float64(X) / Y

	fmt.Printf("%.3f km/l\n", consumption)
}
