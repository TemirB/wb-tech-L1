package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, value := range arr {
		wg.Add(1)
		go func(v int) {
			fmt.Println(v * v)
			defer wg.Done()
		}(value)
	}

	wg.Wait()
}
