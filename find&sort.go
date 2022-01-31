package main

import "fmt"

func FindMaxMini(point *[5]int) (int, int) {
	length := len(*point)
	var mini, max int = point[0], point[0]
	for i := 0; i < length; i++ {
		if mini > point[i] {
			mini = point[i]
		}
		if max < point[i] {
			max = point[i]
		}
	}
	return mini, max
}

func main() {
	var array [5]int = [5]int{3, 2, 4, 5, 1}

	mini, max := FindMaxMini(&array)

	fmt.Println(mini, max)
}
