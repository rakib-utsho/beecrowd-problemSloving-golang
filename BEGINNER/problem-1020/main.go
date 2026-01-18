
// 1020 - Age in Days
package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanln(&N)
	y := N / 365
	r := N % 365

	m := r / 30
	d := r % 30

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", y, m, d)
}
