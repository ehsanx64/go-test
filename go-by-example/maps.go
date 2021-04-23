package main
import "fmt"

func main() {
	// Create an empty map. Following line creates a map where keys are strings
	// and values are ints
	m := make(map[string]int)

	// Set values for map's keys
	m["k1"] = 7
	m["k2"] = 13

	// Dump the map
	fmt.Println("map:", m)

	// Get k1 key value from the map and put it in a variable
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))		// Print the count of our map

	delete(m, "k2")					// Delete k2 key (along with its value) from the map
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map. This can be used to disambiguate between 
	// missing keys and keys with zero values like 0 or "". Here we didn’t need 
	// the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Create a map and initialize it
	n := map[string]int {
		"foo": 1,
		"bar": 2}
	fmt.Println("map:", n)
}