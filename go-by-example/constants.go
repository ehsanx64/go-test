package main

import (
	"fmt"
	"math"
)

// Explicitly define an identifier named 's' as a string constant
const s string = "constant"

func main() {
	// Print out the s
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
