package main

import "fmt"

/*
** Go has built-in support for returning multiple return values from a function.
** It's useful when returning both results and error codes from functions.
 */

/*
** Define a function named 'vals' that returns to ints.
 */
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Get the function retvals and put them in two variables using multiple
	// assigment. Here a contains 3 and b contains 7
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Since vals() returns 2 values; if didn't need one/some of them we use
	// the blank identifier (_) to get unwanted values
	_, c := vals()
	fmt.Println(c)
}
