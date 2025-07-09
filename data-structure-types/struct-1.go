package main

type person struct {
	name string
	age  int
	yob  int
}

func main() {
	p1 := person{
		name: "James",
		age:  25,
		yob:  2000,
	}

	p2 := person{
		name: "charles",
		age:  26,
		yob:  1999,
	}

	println(p1.name, "is of age:", p1.age, "and his birth year is:", p1.yob)
	println(p2.name, "is of age:", p2.age, "and his birth year is:", p2.yob)
}
