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

	nWorkers := 10
	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)

	wg.Add(nWorkers)
	for i := range nWorkers {
		go work(ctx, &wg, i)
	}

	<-sigChan
	fmt.Println("\nПолучен SIGINT, отменяю контекст...")
	cancel()

	wg.Wait()
	fmt.Printf("Все горутины были отмененны")
}

func work(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%d] Поймал отмену контекста\n", i)
			return
		default:
			fmt.Printf("[%d] Работаю\n", i)
			time.Sleep(1 * time.Second)
		}
	}
}
