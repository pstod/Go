package main

import "fmt"

func main() {

	a := []int{}
	for i := 0; i <= 10; i++ {
		a = append(a, i)
	}
	for i := 0; i <= 10; i++ {
		if a[i]%2 == 0 {
			fmt.Println(a[i], "is even")
		} else {
			fmt.Println(a[i], "is odd")
		}
	}

}

/*
➜  Assignment1 git:(master) ✗ go run main.go
0 is even
1 is odd
2 is even
3 is odd
4 is even
5 is odd
6 is even
7 is odd
8 is even
9 is odd
10 is even
*/
