package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Outer loop iteration:", i)
		for j := 0; j < 5; j++ {
			fmt.Println("  Inner loop iteration:", j)
		}
	}
}

// This code demonstrates a nested loop where the outer loop runs 5 times
// and for each iteration of the outer loop, the inner loop runs 5 times.
// The output will show the current iteration of both the outer and inner loops.
// This is useful for scenarios where you need to perform operations on a grid or matrix,
