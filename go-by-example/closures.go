package main

import "fmt"

/*
** This function (intSeq) return an anonymous function. The returned function
** closes over the variable i to form a closure.
 */
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println("1st nextInt:", nextInt())

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
