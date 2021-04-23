package main
import "fmt"

/*
** Define a function named plus which takes two ints as input and return an
** int as output.
*/
func plus(a int, b int) int {
	return a + b
}

/*
** Define a function named plusPlus which takes three ints as input.
** If a group of arguments have the same type, we can omit the type for all
** except the last one. Here a, b and c all are ints.
*/
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}