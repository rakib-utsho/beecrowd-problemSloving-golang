// 1010 - Simple Calculate

package main

import "fmt"

func product() (int, int, float64) {
	var a, b int
	var c float64

	fmt.Scanf("%d %d %f\n", &a, &b, &c)
	return a, b, c
}
func productValue(a int, b int, c float64) float64 {
	var total = float64(b) * c

	return total
}

func totalValue(a float64, b float64) float64 {
	var total = a + b

	return total

}

func main() {
	a, b, c := product()
	total1 := productValue(a, b, c)

	d, e, f := product()
	total2 := productValue(d, e, f)

	result := totalValue(total1, total2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", result)

}
