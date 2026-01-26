package main

import "fmt"

func main() {
	var a float64

	fmt.Scanln(&a)

	if a < 0 || a > 100 {
		fmt.Println("Fora de intervalo")
	} else if a >= 0 && a <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if a >= 25 && a <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if a >= 50 && a <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else {
		fmt.Println("Intervalo (75,100]")
	}
}
