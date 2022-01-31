package main

import "fmt"

func main() {
	// Using var keyboard to declare a variable and assign a value using type inference
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variable at one and initialize them (b and c are ints)
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Using type inference to declare and initialize a boolean value
	var d = true
	fmt.Println(d)

	// Declare variable 'e', but since it is not initialized, it will have default initialization value
	// for ints which is 0
	var e int
	fmt.Println(e)

	// := operator
	f := "apple"
	fmt.Println(f)
}
