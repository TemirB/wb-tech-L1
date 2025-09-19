package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	m  map[string]int
	mu *sync.Mutex
}

func NewSafeMap(m map[string]int, mu *sync.Mutex) *safeMap {
	return &safeMap{
		m:  m,
		mu: mu,
	}
}

func (sm *safeMap) Inc(key string) {
	sm.mu.Lock()
	sm.m[key]++
	sm.mu.Unlock()
}

func (sm *safeMap) Get(key string) int {
	sm.mu.Lock()
	v := sm.m[key]
	sm.mu.Unlock()
	return v
}

func (sm *safeMap) Len() int {
	sm.mu.Lock()
	l := len(sm.m)
	sm.mu.Unlock()
	return l
}

func main() {
	sMap := NewSafeMap(make(map[string]int), &sync.Mutex{})
	var wg sync.WaitGroup

	nKeys := 5
	nWorkers := 10
	opsPerWorker := 100

	wg.Add(nWorkers)
	for w := range nWorkers {
		go func(w int) {
			defer wg.Done()
			for ops := range opsPerWorker {
				key := fmt.Sprintf("k%02d", (w+ops)%nKeys)
				sMap.Inc(key)
			}
		}(w)
	}
	wg.Wait()

	// Проверка суммы
	total := 0
	for k := range nKeys {
		key := fmt.Sprintf("k%02d", k)
		total += sMap.Get(key)
	}
	expected := nWorkers * opsPerWorker

	fmt.Printf("[STATS] len(map)=%d total=%d expected=%d OK=%v\n", sMap.Len(), total, expected, total == expected)
	fmt.Println(sMap.m)
}
