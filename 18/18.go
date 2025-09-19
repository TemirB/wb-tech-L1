package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	v atomic.Uint64
}

func (c *atomicCounter) Inc() {
	c.v.Add(1)
}

func (c *atomicCounter) Load() uint64 {
	return c.v.Load()
}

func main() {
	atomicCount := &atomicCounter{
		v: atomic.Uint64{},
	}

	var wg sync.WaitGroup
	nWorkers := 10
	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000_000; j++ {
				atomicCount.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("counter = %d\n", atomicCount.Load())
}
