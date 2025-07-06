package main

import (
	"fmt"
)

func main() {
	marvel := map[string]int{
		"IronMan":              2008,
		"Thor":                 2011,
		"CaptainAmerica":       2011,
		"Avengers":             2012,
		"GuardiansOfTheGalaxy": 2014,
	}

	// Using the comma-ok idiom to check if a key exists in the map
	if _, ok := marvel["IronMan"]; ok { // This will work because the key is exactly "IronMan"
		fmt.Printf("IronMan is present in the map with release year: %v \n", marvel["IronMan"])
	} else {
		fmt.Println("IronMan is not present in the map")
	}

	if _, ok := marvel["Iron Man"]; ok { // This will not work because the key is not exactly "Iron Man" the way it is defined in the map the key is "IronMan"
		fmt.Println("IronMan is present in the map with release year:", marvel["Iron Man"])
	} else {
		fmt.Println("Iron Man is not present in the map")
	}

	if _, ok := marvel["IronMan3"]; ok { // This will not work because the key "IronMan3" does not exist in the map
		fmt.Println("IronMan3 is present in the map with release year:", marvel["IronMan3"])
	} else {
		fmt.Println("IronMan3 is not present in the map")
	}
}
