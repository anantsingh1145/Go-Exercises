package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	// Using a switch statement to categorize the random number
	// Note: rand.Intn(n) generates a number in the range [0, n), so for 0 to 250, we use rand.Intn(251)
	switch {
	case x <= 100:
		fmt.Printf("Number is between 0 and 100: %v\n", x)
	case x > 100 && x <= 200:
		fmt.Printf("Number is between 101 and 200: %v\n", x)
	case x > 200 && x <= 250:
		fmt.Printf("Number is between 201 and 250: %v\n", x)
	default:
		fmt.Println("Number is out of range")
	}
}
