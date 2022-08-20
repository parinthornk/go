package main

import (
	"fmt"
)

type House struct {
	id      int
	name    string
	members []Person
}

type Phone struct {
	id    int
	name  string
	owner Person
}

func main2() {

	var p1 Person
	var p2 Person
	var p3 Person

	p1.name = "Pa"
	p2.name = "Ja"
	p3.name = "Ta"

	var h1 House
	h1.members = []Person{p1, p2}
	h1.members = append(h1.members, p3)

	//fmt.Println("Hello World!")

	fmt.Println(h1.members[0])

	var ph1 Phone
	ph1.owner = p1
	fmt.Println(ph1)
}
