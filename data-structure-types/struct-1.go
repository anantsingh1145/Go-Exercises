package main

type person struct {
	name string
	age  int
	doy  int
}

func main() {
	p1 := person{
		name: "James",
		age:  25,
		doy:  2000,
	}

	p2 := person{
		name: "charles",
		age:  26,
		doy:  1999,
	}

	println(p1.name, "is of age:", p1.age, "and his birth year is:", p1.doy)
	println(p2.name, "is of age:", p2.age, "and his birth year is:", p2.doy)
}
