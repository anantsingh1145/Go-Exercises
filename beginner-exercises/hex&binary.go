package main

import (
	"fmt"
)

func main() {
	a := 44
	fmt.Printf("%v \t %T \t %b \t %o \t %#x \n", a, a, a, a, a) // Print value, type, binary, octal, and hexadecimal representations of 'a'

	// Hexadecimal and binary representations of numbers
	hexValue := 0x1A       // Hexadecimal representation of 26
	binaryValue := 0b11010 // Binary representation of 26 (Go supports binary literals since version 1.13)

	fmt.Printf("Hexadecimal value: %X\n", hexValue) // %X formats the number as hexadecimal
	fmt.Printf("Binary value: %b\n", binaryValue)   // %b formats the number as binary

	// Converting decimal to hexadecimal and binary
	decimalValue := 26
	fmt.Printf("Decimal %d in hexadecimal is %X\n", decimalValue, decimalValue)
	fmt.Printf("Decimal %d in binary is %b\n", decimalValue, decimalValue)
}
