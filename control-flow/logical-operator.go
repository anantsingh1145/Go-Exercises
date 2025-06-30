package main

import (
	"fmt"
)

func main() {

	// Logical Operators
	// && (AND)
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && true)  // false
	fmt.Println(false && false) // false
	fmt.Println("---------------")

	//|| (OR), ! (NOT)
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || true)  // true
	fmt.Println(false || false) // false
	fmt.Println("---------------")

	// ! (NOT)
	fmt.Println(!true)  // false
	fmt.Println(!false) // true

	fmt.Println("---------------")
	fmt.Println("---------------")
	// Short-circuit evaluation
	fmt.Println("true && (5 > 3): ", true && (5 > 3))   // true, second condition evaluated
	fmt.Println("false && (5 > 3): ", false && (5 > 3)) // false, second condition not evaluated
	fmt.Println("true || (5 < 3): ", true || (5 < 3))   // true, second condition not evaluated
	fmt.Println("false || (5 < 3): ", false || (5 < 3)) // false, second condition evaluated
	fmt.Println("---------------")

	fmt.Println("Logical operators are used to combine or invert boolean values.")
	fmt.Println("They are commonly used in conditional statements and loops to control the flow of execution.")

}
