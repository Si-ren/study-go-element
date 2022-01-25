package main

import (
	"fmt"
	"os"
)

func main() {
	slice := make([]string, 5)
	slice[1] = "AAAAA"
	slice[2] = "BBBBB"
	slice[3] = "cccccc"
	slice[4] = "ddddd"

	test := slice[2:3:3]
	test = append(test, "11111")
	test[0] = "test"

	fmt.Println(len(slice))
	fmt.Println(len(test))
	fmt.Println(cap(slice))
	fmt.Println(cap(test))

	/*
		for _, value := range slice{
			fmt.Printf("Value: %s\n", value)
		}
		for _, value := range test{
			fmt.Printf("Value: %s\n", value)
		}
	*/
	os.Exit(2)
}
