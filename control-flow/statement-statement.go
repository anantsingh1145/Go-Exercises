package main

import (
	"fmt"
)

func main() {
	m := map[string]int{ // Initializing a map with string keys and integer values
		"maven":  28,
		"gradle": 30,
		"sbt":    32,
		"ant":    25,
	}

	if s := m["gradle"]; m["gradle"] > 28 {
		if m["gradle"] < 32 {
			fmt.Println("Gradle is elder than maven and younger than sbt as his age is:", s) // If the value associated with the key "gradle" is greater than 28, print it.}
		}
	}
}

// If the value associated with the key "gradle" is greater than 28 and less than 32, print a message indicating that Gradle is older than Maven and younger than SBT
