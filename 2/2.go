package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	// Squaring
	for _, value := range numbers {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v * v)
			defer wg.Done()
		}(value)
	}

	wg.Wait()
}
