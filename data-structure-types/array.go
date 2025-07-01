package main

import (
	"fmt"
)

// This program demonstrates the use of an array in Go.
// It creates an array of movie titles and prints the array and its length.
// Arrays in Go are fixed-size collections of elements of the same type.
// The length of an array is determined at compile time and cannot be changed.
// The program initializes an array with movie titles and prints the array and its length.
// Arrays are useful when you know the number of elements in advance and want to ensure that the
// size remains constant throughout the program.
func main() {
	a := [...]string{"Iron Man", "The Avengers", "Tropic Thunder", "Sherlock Holmes", "Kiss Kiss, Bang Bang", "Zodiac", "Chaplin", "Good Night, and Good Luck", "The Judge", "Oppenheimer"}

	fmt.Println("Best movies of Robert Downey Jr. :", a)
	fmt.Println("Length of the array:", len(a))
}
