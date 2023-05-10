package main

import "fmt"

type numstr []int

func main() {
	a := numstr{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	create_nstr(a) //Define,initialize and auto assign data type

}
