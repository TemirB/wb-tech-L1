package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	base := arr[len(arr)/2]
	var left, middle, right []int

	for _, item := range arr {
		if item < base {
			left = append(left, item)
		} else if item == base {
			middle = append(middle, item)
		} else {
			right = append(right, item)
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	result := append(append(sortedLeft, middle...), sortedRight...)
	return result
}

func main() {
	testCases := [][]int{
		{64, 34, 25, 12, 22, 11, 90},
		{5, 2, 3, 1, 4},
		{1},
		{},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{2, 3, 1},
		{3, 1, 2},
	}

	for i, arr := range testCases {
		fmt.Printf("Тест %d:\n", i+1)
		fmt.Printf("  До: %v\n", arr)
		sorted := quickSort(arr)
		fmt.Printf("  После: %v\n\n", sorted)
	}
}
