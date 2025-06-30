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

	if v, ok := m["sbt"]; ok { // Using the comma-ok idiom to check if the key "sbt" exists in the map // If the key exists, ok will be true and k will hold the value associated with the key "sbt"
		// If the key does not exist, ok will be false and k will be the zero value for the map's value type (int in this case)
		fmt.Println("Age of sbt:", v) // If the key "sbt" exists, print its value
	} else if v, ok := m["xyz"]; ok {
		fmt.Println("Age of xyz:", v) // If the key "xyz" exists, print its value
	} else {
		fmt.Println("Key not found") // If neither key exists, print a message
	}
}
