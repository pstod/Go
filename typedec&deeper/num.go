package main

import (
	"fmt"
	"math/rand"
)

func create_nstr(a numstr) {

	for i := range a {
		b := rand.Intn(len(a) - 1)
		a[i], a[b] = a[b], a[i]
		fmt.Println(a)
	}

}

/*➜  typedec&deeper git:(master) ✗ go run main.go num.go
[0 1 2 3 4 5 6 7 8 9]
[2 1 0 3 4 5 6 7 8 9]
[2 7 0 3 4 5 6 1 8 9]
[0 7 2 3 4 5 6 1 8 9]
[0 7 2 4 3 5 6 1 8 9]
[0 7 2 4 6 5 3 1 8 9]
[5 7 2 4 6 0 3 1 8 9]
[5 3 2 4 6 0 7 1 8 9]
[5 3 2 4 6 0 7 1 8 9]
[8 3 2 4 6 0 7 1 5 9]
[8 3 2 4 6 0 7 1 9 5]*/
