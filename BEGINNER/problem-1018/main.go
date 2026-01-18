// beecrowd | 1018 Banknotes
package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	original := n

	notes := []int{100, 50, 20, 10, 5, 2, 1}

	fmt.Println(original)

	for _, d := range notes {
		count := n / d
		n = n % d

		fmt.Printf("%d nota(s) de R$ %d,00\n", count, d)
	}
}
