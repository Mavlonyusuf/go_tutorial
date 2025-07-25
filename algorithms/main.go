package main

import "fmt"

func main() {
	fmt.Println(binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}
func binary_search(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return guess
		} else if guess > item {
			high = mid - 1
		} else {
			high = mid + 1
		}
	}
	return -1
}