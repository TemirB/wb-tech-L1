package main

import "fmt"

func deleteAt[T any](s []T, i int) ([]T, bool) {
	if i < 0 || i >= len(s) {
		return s, false
	}
	// копируем слайс со сдвигом
	copy(s[i:], s[i+1:])

	// тут подменяем последний элемент nil-ом чтобы GC подчистил память
	var nilValue T
	s[len(s)-1] = nilValue

	// режем длину слайса
	return s[:len(s)-1], true
}

func main() {
	sourceSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// idx 				 0  1  2  3  4  5  6  7  8  9

	fmt.Printf("Enter index between 0 and %d\n", len(sourceSlice)-1)
	var i int
	fmt.Scan(&i)

	res, ok := deleteAt(sourceSlice, i)
	if !ok {
		fmt.Println("index out of range")
		return
	}
	fmt.Printf("source after delete: %v\n", res)
}
