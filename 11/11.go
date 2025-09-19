package main

import "fmt"

func Contain(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	result := []int{}
	for _, t := range B {
		if Contain(A, t) {
			result = append(result, t)
		}
	}

	fmt.Println(result)
}
