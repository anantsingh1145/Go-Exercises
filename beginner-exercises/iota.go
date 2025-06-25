package main

import (
	"fmt"
)

const (
	_  = iota // 0
	x1        // 1
	x2        // 2
	x3        // 3
	x4        // 4
	x5        // 5
	x6        // 6
)

func main() {
	fmt.Println(x1, x2, x3, x4, x5) // Print the values of the constants
	// Iota is a predeclared identifier in Go that simplifies constant declarations.
	fmt.Printf("%d \t %b\n", 1, 1)         // Print the decimal and binary representation of 1
	fmt.Printf("%d \t %b\n", 1<<x1, 1<<x1) // Print the decimal and binary representation of 1 shifted left by 1 bits
	fmt.Printf("%d \t %b\n", 1<<x2, 1<<x2) // Print the decimal and binary representation of 1 shifted left by 2 bits
	fmt.Printf("%d \t %b\n", 1<<x3, 1<<x3) // Print the decimal and binary representation of 1 shifted left by 3 bits
	fmt.Printf("%d \t %b\n", 1<<x4, 1<<x4) // Print the decimal and binary representation of 1 shifted left by 4 bits
	fmt.Printf("%d \t %b\n", 1<<x5, 1<<x5) // Print the decimal and binary representation of 1 shifted left by 5 bits
	fmt.Printf("%d \t %b\n", 1<<x6, 1<<x6) // Print the decimal and binary representation of 1 shifted left by 6 bits
}
