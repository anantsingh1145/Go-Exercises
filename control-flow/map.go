package main

import (
	"fmt"
)

// This program demonstrates the use of a map in Go.
// It creates a map with string keys and integer values, iterates over the map,
// and prints each key-value pair along with the total number of elements in the map.
// Maps in Go are unordered collections of key-value pairs.
func main() {
	m := map[string]int{ // Initializing a map with string keys and integer values
		"maven":  28,
		"gradle": 30,
		"sbt":    32,
		"ant":    25,
	}

	for k, v := range m { // Using a for loop with range to iterate over the map
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	fmt.Println("Total number of elements in the map:", len(m)) // Printing the total number of elements in the map
}
