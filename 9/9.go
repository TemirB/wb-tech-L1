package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)

	// Для корректного завершения main горутины
	var wg sync.WaitGroup
	wg.Add(2)

	// Буферизация, для уменьшения блокировок
	numbers := make(chan int, 16)
	squares := make(chan int, 16)

	// Горутина, для отмены контекста
	go func() {
		sig := <-sigChan
		fmt.Printf("\nReceived signal: %v, shutting down...\n", sig)
		cancel()
	}()

	// Генератор значений
	go func() {
		defer close(numbers)
		defer wg.Done()
		x := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("generator: graceful shutdown")
				return
			default:
				select {
				case numbers <- x:
					x++
				case <-ctx.Done():
					return
				}
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Вычислитель квадратов
	go func() {
		defer close(squares)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("squarer: graceful shutdown")
				return
			case val, ok := <-numbers:
				if !ok {
					return
				}
				select {
				case squares <- val * val:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("main: graceful shutdown completed")
			return
		case squared, ok := <-squares:
			if !ok {
				fmt.Println("main: channel closed, shutting down")
				return
			}
			fmt.Printf("squared value: %d\n", squared)
		}
	}
}
