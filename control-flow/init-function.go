package main

import (
	"fmt"
	"math/rand"
)

// init function is called before main
// It is often used for initialization tasks such as setting up variables, logging, or other setup
func init() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is: ", rand.Intn(1000)) // Simulating a time output
}

func main() {
	x := rand.Intn(250)

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
