package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	yob  int
}

func main() {
	p1 := struct { // This is how we define the Anonymous struct where we just need to restrict the use of struct in some functions or a set of code.
		name string
		age  int
		yob  int
	}{
		name: "James",
		age:  25,
		yob:  2000,
	}

	p2 := person{
		name: "charles",
		age:  26,
		yob:  1999,
	}

	fmt.Println(p1.name, "is of age:", p1.age, "and his birth year is:", p1.yob)
	fmt.Println(p2.name, "is of age:", p2.age, "and his birth year is:", p2.yob)
}
