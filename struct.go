package main

import "fmt"

const (
	k1 string = "test1"
	k2 int    = 1
)

func main() {
	fmt.Println(k1)
	fmt.Println(k2)

	type user struct {
		name  string
		age   int
		alive bool
	}

	tt := user{"aaaa", 8, true}
	fmt.Println(tt)
}
