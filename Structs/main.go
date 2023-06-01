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
	fmt.Println()
	p.updatename("lol") //shortcut : go converts person to &person due to receiver type
	fmt.Println(p)
	(&p).updatename("mol")
	fmt.Println(p)
}

func (p person) printout() {

	fmt.Println(p)
	fmt.Printf("%+v", p)
}

func (p *person) updatename(b string) {

	(*p).name = b

}
