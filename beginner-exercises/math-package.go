package main

import (
	"fmt"
	"math"
)

func mathematics() {
	// Basic arithmetic operations
	a := 10
	b := 20

	sum := a + b
	diff := a - b
	prod := a * b
	quot := b / a

	fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %d\n", sum, diff, prod, quot)

	// Using math package for advanced calculations
	sqrtValue := math.Sqrt(16)
	powValue := math.Pow(2, 3)

	fmt.Printf("Square root of 16: %.2f, 2 raised to the power of 3: %.2f\n", sqrtValue, powValue)
}

func main() {
	fmt.Println("Welcome to the math package example!")

	// Call the math function to demonstrate basic arithmetic and advanced calculations
	mathematics()

	// Additional examples of using math functions
	fmt.Printf("Pi: %.2f\n", math.Pi)                               // Print the value of Pi
	fmt.Printf("Cosine of 0 radians: %.2f\n", math.Cos(0))          // Cosine of 0 radians
	fmt.Printf("Sine of Pi/2 radians: %.2f\n", math.Sin(math.Pi/2)) // Sine of Pi/2 radians
}

// Note: The math package is part of the Go standard library and provides basic constants and mathematical functions.
// To use it, you need to import the "math" package at the beginning of your Go file.
// Make sure to run `go get math` if you are using a module system, although the math package is included in the standard library.
// The above code demonstrates basic arithmetic operations and advanced calculations using the math package in Go.
// The math package provides functions for basic constants and mathematical operations.
// The example includes basic arithmetic operations, square root, power calculations, and trigonometric functions.
// The output will show the results of these calculations in a formatted manner.
// The math package is a standard library in Go, so you don't need to install it separately.
// The code is structured to demonstrate the use of the math package in a simple Go program.
