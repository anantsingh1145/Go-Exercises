package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)

	// Using if-else statements to check conditions on x and y
	if x < 4 && y < 4 {
		fmt.Printf("Both x = %d and y = %d are less than 4\n", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("Both x = %d and y = %d are less than 6\n", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x = %d is between 4 and 6, and y = %d is less than or equal to 6\n", x, y)
	} else if y != 5 {
		fmt.Printf("y = %d is not equal to 5\n", y)
	} else {
		fmt.Printf("No specific condition met for x = %d and y = %d \n", x, y)
	}
}
