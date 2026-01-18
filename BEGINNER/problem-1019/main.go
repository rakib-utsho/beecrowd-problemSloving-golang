
// 1019 - Time Conversion
package main

import "fmt"

func main() {
	var seconds, hours, minute int

	fmt.Scanln(&seconds)

	hours = seconds / 3600
	seconds = seconds - (hours * 3600)

	minute = seconds / 60
	seconds = seconds - (minute * 60)

	fmt.Printf("%d:%d:%d\n", hours, minute, seconds)

}
