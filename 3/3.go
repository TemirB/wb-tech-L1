package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	var nWorkers int
	fmt.Printf("Enter number of workers\n")
	if _, err := fmt.Scan(&nWorkers); err != nil || nWorkers <= 0 {
		fmt.Printf("Error while scan, or number of workers is negative")
		return
	}

	wg.Add(nWorkers)
	for i := range nWorkers {
		go func(id int) {
			defer wg.Done()
			for data := range ch {
				fmt.Printf("[%d] %d\n", id, data)
			}
		}(i)
	}

	for i := 1; ; i++ {
		ch <- i
		time.Sleep(50 * time.Millisecond)
	}
}
