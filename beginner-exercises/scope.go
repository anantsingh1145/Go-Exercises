package main

import (
	"fmt"
)

var var1 int = 10 // Global variable

const const1 float64 = 20 // Global constant

func main() {
	var2 := float64(var1)

	result := var2 * const1

	fmt.Printf("The result of multiplying var2 and const1 is: %.2f\n", result)
}
