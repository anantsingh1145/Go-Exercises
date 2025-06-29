package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is first line for sequential flow program")
	fmt.Println("This is second line for sequential flow program")

	x := rand.Intn(250) // Generates a random number between 0 and 249
	// Note: rand.Intn(n) generates a number in the range [0, n), so for 0 to 250, we use rand.Intn(251)
	// If you want to include 250, use rand.Intn(251)
	// rand.Seed(time.Now().UnixNano()) // Uncomment this line to seed the random number generator for different results each time
	// Note: Seeding is not necessary for this example, but it's good practice if you want different random numbers each time you run the program.

	fmt.Printf("Random number generated is: %v\n", x) // Print the generated random number

	// Conditional statements to check the range of the generated number
	if x <= 100 {
		fmt.Printf("Number is between 0 and 100: %v\n", x)
	}

	if x > 100 && x <= 200 {
		fmt.Printf("Number is between 101 and 200: %v\n", x)
	}
	if x > 200 && x <= 250 {
		fmt.Printf("Number is between 201 and 250: %v\n", x)
	}
}
