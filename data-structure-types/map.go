package main

import (
	"fmt"
)

// This program demonstrates the use of a map in Go.
// It creates a map to store Marvel movies and their release years.
// Maps in Go are unordered collections of key-value pairs.
// The keys in this map are movie titles, and the values are their respective release years.
// The program prints the length of the map and its contents.
// Maps are useful for associating unique keys with values, allowing for efficient lookups and retrieval
func main() {

	marvel := map[string]int{
		"Iron-man":                2008,
		"Thor":                    2011,
		"Captain America":         2011,
		"Avengers":                2012,
		"Guardians of the Galaxy": 2014,
	}

	fmt.Println("Length of the map is:", len(marvel))
	fmt.Printf("Marvel movies and their release years: %v", marvel)
}
