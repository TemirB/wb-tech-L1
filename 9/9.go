package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)

	xCh := make(chan int)
	xSquaredCh := make(chan int)

	go func() {
		sig := <-sigChan
		fmt.Printf("\nReceived signal: %v, shutting down...\n", sig)
		cancel()
	}()

	// Генератор значений
	go func() {
		defer close(xCh)
		x := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("generator: graceful shutdown")
				return
			default:
				select {
				case xCh <- x:
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
		defer close(xSquaredCh)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("squarer: graceful shutdown")
				return
			case val, ok := <-xCh:
				if !ok {
					return
				}
				select {
				case xSquaredCh <- val * val:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("main: graceful shutdown completed")
			time.Sleep(100 * time.Millisecond)
			return
		case squared, ok := <-xSquaredCh:
			if !ok {
				fmt.Println("main: channel closed, shutting down")
				return
			}
			fmt.Printf("squared value: %d\n", squared)
		}
	}
}
