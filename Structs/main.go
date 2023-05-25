package main

import "fmt"

type address struct {
	city string
	zip  int
}

type person struct {
	name string
	age  int
	address
}

func main() {

	p := person{"Pd", 23, address{"Pune", 90908}}
	p.printout()
}

func (p person) printout() {

	fmt.Println(p)
}
