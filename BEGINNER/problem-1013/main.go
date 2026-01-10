//1013 - The Greatest

package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d\n", &a, &b, &c)

	var maiorAB = (a + b + abs(a-b)) / 2

	var gretestABC = (maiorAB + c + abs(maiorAB-c)) / 2

	fmt.Println(gretestABC, "eh o maior")

}
