package main

import (
	"fmt"
)

func accumulation(initialNumber int) func(int) int {
	return func(num int) int {
		for i := 0; i <= num; i++ {
			initialNumber += i
		}
		return initialNumber
	}

}

func main() {
	f := accumulation(10)
	fmt.Printf("consequence is %d \n", f(5))
	fmt.Printf("consequence is %d \n", f(3))
}
