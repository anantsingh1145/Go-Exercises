package main

import (
	"fmt"
)

const (
	x1 = iota // 0
	x2        // 1
	x3        // 2
	x4        // 3
	x5        // 4
)

func main() {
	fmt.Println(x1, x2, x3, x4, x5) // Print the values of the constants
	// Iota is a predeclared identifier in Go that simplifies constant declarations.
	fmt.Printf("%d \t %b\n", 1, 1) // Print the decimal and binary representation of 1
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2) // Print the decimal and binary representation of 1 shifted left by 2 bits
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3) // Print the decimal and binary representation of 1 shifted left by 3 bits
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4) // Print the decimal and binary representation of 1 shifted left by 4 bits
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5) // Print the decimal and binary representation of 1 shifted left by 5 bits
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6) // Print the decimal and binary representation of 1 shifted left by 6 bits
}
