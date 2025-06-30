package main

import (
	"fmt"
)

// This program demonstrates how to iterate over a slice and print both the index and value of each element.
// It uses a for loop with the range keyword to achieve this.
func main() {
	xi := []int{10, 20, 30, 40, 50}
	for i, v := range xi {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	// Using a for loop to iterate over the slice and print index and value
}
