package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// Using a loop to generate random numbers and check conditions
	for i := 0; i < 100; i++ {
		x, y := rand.Intn(10), rand.Intn(10)
		fmt.Printf("Iteration %v:\t", i)

		// Using a switch statement to check conditions on x and y
		switch {
		case x < 4 && y < 4:
			fmt.Printf("Both x = %d and y = %d are less than 4\n", x, y)
		case x > 6 && y > 6:
			fmt.Printf("Both x = %d and y = %d are greater than 6\n", x, y)
		case x >= 4 && x <= 6:
			fmt.Printf("x = %d is between 4 and 6, and y = %d is less than or equal to 6\n", x, y)
		case y != 5:
			fmt.Printf("y = %d is not equal to 5\n", y)
		default:
			fmt.Printf("No specific condition met for x = %d and y = %d \n", x, y)
		}
	}
}
