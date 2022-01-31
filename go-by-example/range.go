package main

import "fmt"

func main() {
	// Range iterates over elements in a variety of data structures

	/*
	** Using range to sum the numbers in a slice
	 */
	nums := []int{2, 3, 4}
	sum := 0

	//
	fmt.Println("Iterate through range and display both index and value")
	fmt.Println("Legend:")
	fmt.Println("index -> value")
	for i, v := range nums {
		fmt.Printf("%d -> %d\n", i, v)
	}

	// When using range on arrays and slices it return both the key (index) and value
	// Here since with did not need the index we ignore it using _ identifier
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// print the index which hold the value 3
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range over maps iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	// Iterate over keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points (characters). The
	// first value (i) is the index and the second (c) is the value.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
