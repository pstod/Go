package main

import "fmt"

type abook map[string]string

func main() {

	addressbk := map[string]string{"SD": "Kharagpur"}
	addressbk["PD"] = "Pune"
	addressbk["VD"] = "Mumbai"
	fmt.Println(addressbk)

	a := abook{"ABC": "CBA"}

	a.printout()
}

func (p abook) printout() {
	fmt.Println(p)
}

/*
➜  maps git:(master) ✗ go run main.go
map[PD:Pune SD:Kharagpur VD:Mumbai]
map[ABC:CBA]
*/
