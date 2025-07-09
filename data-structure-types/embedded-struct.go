package main

import "fmt"

// This program demonstrates how to use embedded structs in Go.
// It defines a `person` struct and a `secretAgent` struct that embeds `person
type person struct {
	name string
	age  int
	yob  int
}

// The `secretAgent` struct has an additional field `licenseToKill` and
// it also has a `name` field that will overwrite the `name` field in the person struct.
type secretAgent struct {
	person
	licenseToKill bool
	name          string // This will overwrite the name field in person struct
}

func main() {
	p1 := secretAgent{
		person: person{
			name: "James",
			age:  25,
			yob:  2000,
		},
		licenseToKill: true,
		name:          "Stephen",
	}

	p2 := person{
		name: "charles",
		age:  26,
		yob:  1999,
	}

	fmt.Println(p1.person.name, "is of age:", p1.age, "and his birth year is:", p1.yob)
	fmt.Println("Secret Agent:", p1.name, "has license to kill:", p1.licenseToKill)
	fmt.Println(p2.name, "is of age:", p2.age, "and his birth year is:", p2.yob)

}
