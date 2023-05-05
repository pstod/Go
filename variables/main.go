package main

import "fmt"

var b string //assignment not allowed outside func for var

func main() {
	a := 9 //Define,initialize and auto assign data type
	fmt.Println("Variable a = ", a)

	b = "HI"
	fmt.Println(b)
}
