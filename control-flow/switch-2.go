package main

import (
	"fmt"

	"math/rand"
)

func main() {
	for i := 0; i <= 42; i++ {
		x := rand.Intn(5) // Generate a random number between 0 and 5

		fmt.Printf("Iteration %v:\t", i)
		switch x {
		case 0:
			fmt.Printf("Value of x is %v \n", x)
		case 1:
			fmt.Printf("Value of x is %v \n", x)
		case 2:
			fmt.Printf("Value of x is %v \n", x)
		case 3:
			fmt.Printf("Value of x is %v \n", x)
		case 4:
			fmt.Printf("Value of x is %v \n", x)
		default:
			fmt.Printf("Value of x is out of range")
		}
	}
}
