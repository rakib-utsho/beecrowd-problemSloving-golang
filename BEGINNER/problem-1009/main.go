// 1009 - Salary with Bonus

package main

import (
	"fmt"
)

func main() {
	var name string
	var salary, sales float64

	fmt.Scan(&name)
	fmt.Scan(&salary)
	fmt.Scan(&sales)

	total := salary + sales*0.15

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
