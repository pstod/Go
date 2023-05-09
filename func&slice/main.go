package main

import "fmt"

func main() {
	a := []string{"spades", "ace", "clubs", cname()} //Define,initialize and auto assign data type

	for index, ele := range a {
		fmt.Println("value at ", index, " ", ele)
	}
}

func cname() string {

	return "hearts"

}

/* ➜  func&slice git:(master) ✗ go run main.go
value at  0   spades
value at  1   ace
value at  2   clubs
value at  3   hearts */
//slice can have one type of elements
