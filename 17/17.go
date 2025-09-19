package main

import "fmt"

func bSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Избегаем переполнения

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	testcases := [][]int{
		{1, 2, 3},
		{1, 2, 3, 4, 5},
		{1, 3, 5, 7, 9},
		{},
		{1},
	}

	targets := []int{3, 6, 1, 0, 1}

	for idx, test := range testcases {
		target := targets[idx]
		result := bSearch(test, target)
		fmt.Printf("Test %d: arr=%v, target=%d, result=%d\n",
			idx+1, test, target, result)
	}
}
