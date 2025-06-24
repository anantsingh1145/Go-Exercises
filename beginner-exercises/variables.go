package main

import (
	"fmt"
)

func main() {

	a := 40 // ":=" operator is used for short variable declaration. It's a concise way to declare and initialize a variable in one step without explicitly stating its type.
	println("Value of a:", a)

	b, c, d, _, f := 10, 20, 30, 80, "Starting" // Multiple variables can be declared and initialized in a single line.
	// The underscore (_) is used as a blank identifier to ignore a value.
	// It is often used when you want to ignore a value returned by a function or when you don't need a particular variable.
	// In this case, the value 80 is ignored and not assigned to any variable.
	// The variables b, c, d, and f are initialized with the values 10, 20, 30, and "Starting" respectively.
	// The variables b, c, d are of type int and f is of type string
	// The short variable declaration operator ":=" is used to declare and initialize these variables.
	println("Values of b, c, d, f:", b, c, d, f)

	var g int
	println("Value of g (uninitialized):", g)    // The variable g is declared but not initialized, so it defaults to 0.
	g = 50                                       // Now g is assigned a value of 50.
	println("Value of g (after assignment):", g) // Now g has a value of 50.

	// Declare and initialize variables
	var name string = "John"
	var age int = 30
	var height float64 = 5.9
	var isEmployed bool = true

	// Print the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Employed:", isEmployed)

	// Print a formatted string using the variables
	fmt.Printf("%s is %d years old, %.1f feet tall, and employed: %t\n", name, age, height, isEmployed)
}
