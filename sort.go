package main

import (
	"fmt"
)

func bubbleSort(point *[5]int) *[5]int {
	length := len(*point) - 1
	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if point[j] > point[j+1] {
				point[j], point[j+1] = point[j+1], point[j]
			}
		}
	}

	return point
}

func main() {
	var array [5]int = [5]int{3, 2, 4, 5, 1}
	bubbleSort(&array)
	fmt.Println(array)
}
