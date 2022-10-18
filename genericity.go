package main

import "fmt"

type Type01 interface {
	int | float32
}

type Type02 interface {
	int | float32
}

func Add[T int | float32](i1 T, i2 T) float32 {
	return float32(i1) + float32(i2)
}

func main() {
	var a, b float32 = 1.1, 2.2
	fmt.Println(Add(a, b))
}
