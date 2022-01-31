package main

import "fmt"

/*
** A variadic function is a function which its arguments are variable. A good
** example is the fmt.Println() function.
 */

/*
** sum() function takes an arbitrary number of int and return their sum.
 */
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// If we needed to (for example) feed an array to a variadic function.
	// array name must be prefixed with '...' (to unpack the array I guess)
	sum(nums...)
}
