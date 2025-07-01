package main

import (
	"fmt"
)

func main() {
	s := []string{"Iron Man", "The Avengers", "Tropic Thunder", "Sherlock Holmes", "Kiss Kiss, Bang Bang", "Zodiac", "Chaplin", "Good Night, and Good Luck", "The Judge", "Oppenheimer"}
	fmt.Println("Best movies of Robert Downey Jr. :", s)

	for i, v := range s[:3] { // Slicing the slice to only include the first 3 movies
		// The range loop iterates over the first 3 elements of the slice s
		fmt.Println("Index no", i, "and it's Value:", v)
	}

	for i, v := range s[3:6] { // Slicing the slice to include the next 3 movies
		// The range loop iterates over the next 3 elements of the slice s
		fmt.Println("Index no", i, "and it's Value:", v)
	}

	for i, v := range s[6:] { // Slicing the slice to include the remaining movies
		// The range loop iterates over the remaining elements of the slice s
		// This will include all movies from index 7 to the end of the slice
		//
		fmt.Println("Index no", i, "and it's Value:", v)
	}

	fmt.Println("--------------------------------------------")

	for i, v := range s { // This loop iterates over all elements of the slice s
		// The range loop iterates over all elements of the slice s
		// It will print each movie title along with its index
		fmt.Println("Index no", i, "and it's Value:", v)
	}
}
