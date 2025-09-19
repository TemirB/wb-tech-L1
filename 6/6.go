package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 1) выход по условию
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			if i == 0 {
				fmt.Printf("1) Завершили горутину, по условию i = %d\n", i)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 2) Выход по каналу
	cancelCh := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-cancelCh:
				fmt.Printf("2) Завершили горутину, по каналу\n")
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(200 * time.Millisecond)
	cancelCh <- 1

	// 3) Контекст
	ctxWC, cancelWC := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctxWC.Done():
				fmt.Printf("3) Поймали отмену контекста\n")
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(200 * time.Millisecond)
	cancelWC()

	// 4) runtime.Goexit()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			fmt.Printf("4) Завершили горутину через runtime.Goexit()\n")
		}()

		time.Sleep(100 * time.Millisecond)
		runtime.Goexit()
	}()

	// 5) Паника (с восстановлением)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("5) Восстановились после паники: %v\n", r)
			}
		}()

		time.Sleep(100 * time.Millisecond)
		panic("искусственная паника")
	}()

	// 6) Закрытие канала
	dataCh := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			fmt.Printf("6) Завершили горутину по закрытию канала\n")
		}()

		for {
			_, ok := <-dataCh
			if !ok {
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(300 * time.Millisecond)
	close(dataCh)

	wg.Wait()
}
