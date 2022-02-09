package main

import (
	"fmt"
)

func BubbleSort(point *[5]int) {
	length := len(*point) - 1
	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if point[j] > point[j+1] {
				point[j], point[j+1] = point[j+1], point[j]
			}
		}
	}

}

func QuickSort(arr *[5]int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr *[5]int, left, right int) {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
}
func swap(arr *[5]int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr *[5]int, left, right int) int {
	index := left + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[left] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, left, index-1)
	return index - 1
}

func main() {
	var array01 [5]int = [5]int{3, 2, 4, 5, 1}
	BubbleSort(&array01)
	fmt.Println(array01)

	var array02 [5]int = [5]int{3, 2, 4, 5, 1}
	QuickSort(&array02)
	fmt.Println(array02)
}
