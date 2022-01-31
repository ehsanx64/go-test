package main

import "fmt"

func main() {
	/*
	** Define an array of ints with 5 elements. As default value for int's is 0.
	** This array is an array of 5 zero's
	 */
	var a [5]int

	// Println will dump entire array
	fmt.Println("emp:", a)

	// Put 100 in fourth element of 'a' array
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len function return array count
	fmt.Println("len:", len(a))

	// Using braces default values for an array can be specified
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Define a two-dimensional array
	var twoD [2][3]int

	// Using two nested loops set values for twoD array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
