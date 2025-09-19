package main

import "fmt"

func main() {
	aValue := 6
	bValue := -6

	fmt.Printf("a = %d \t b = %d\n", aValue, bValue)

	aValue += bValue
	bValue = aValue - bValue
	aValue = aValue - bValue

	fmt.Printf("a = %d \t b = %d\n", aValue, bValue)
}
