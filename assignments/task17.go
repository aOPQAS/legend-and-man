package main

import "fmt"

func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1

	for left <= right {
		mid := left + (right-left)/2

		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	item := []int{1, 5, 8, 9, 67, 114, 785}
	target := 63
	result := binarySearch(item, target)

	if result != -1 {
		fmt.Printf("%d", result)
	} else {
		fmt.Println("Element not found.")
	}
}
