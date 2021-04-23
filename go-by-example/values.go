package main

import "fmt"

func main() {
	// Strings can be concatenated using + operator
	fmt.Println("go" + "lang")

	// + adds integers
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("7.0 / 3.0 = ", 7.0 / 3.0)

	// Dumping booleans using fmt.Println()
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}