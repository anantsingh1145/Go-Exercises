package main

import (
	"fmt"
)

// This program demonstrates the use of a map in Go ans using for range loop with the map.
func main() {
	marvel := map[string]int{
		"Iron man":                2008,
		"Thor":                    2011,
		"Captain America":         2011,
		"Avengers":                2012,
		"Guardians of the Galaxy": 2014,
	}

	for k, v := range marvel { // The range loop iterates over the map marvel
		// k is the key (movie title) and v is the value (release year)
		fmt.Println("Movie", k, "was released in the year:", v)
	}

	for _, v := range marvel { // This loop iterates over the values of the map marvel
		// It prints only the release years of the movies
		fmt.Println("Release year:", v)
	}
	for k := range marvel { // This loop iterates over the keys of the map marvel
		// It prints only the movie titles
		fmt.Println("Movie title:", k)
	}
}
