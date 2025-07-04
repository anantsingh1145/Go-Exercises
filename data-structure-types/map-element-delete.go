package main

import (
	"fmt"
)

func main() {
	marvel := map[string]int{
		"Iron man":                2008,
		"Thor":                    2011,
		"Captain America":         2011,
		"Avengers":                2012,
		"Guardians of the Galaxy": 2014,
	}
	fmt.Printf("Marvel movies and their release years: %v \n", marvel)

	delete(marvel, "Guardians of the Galaxy") // This line deletes the key "Guardians of the Galaxy" and its associated value from the map marvel

	delete(marvel, "Guardinans of the Galaxy") // This line attempts to delete a key that does not exist in the map, which has no effect
	fmt.Println("After deletion, Guardinans of the Galaxy which is not present in map data:", marvel)
	fmt.Println("Printing the wrong name of the movie, which is not present in the map data:", marvel["Guradiansas of the Galaxy"]) // This line will cause a runtime panic because the key does not exist in the map

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
