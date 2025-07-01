package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d", "e", "f"}

	fmt.Println("slice1: ", slice1) // Output: [a b c]
	fmt.Println("slice2: ", slice2) // Output: [d e f]

	s1Xs2 := [][]string{slice1, slice2}
	fmt.Println("slice1 X slice2: ", s1Xs2) // Output: [[a b c] [d e f]]
}
