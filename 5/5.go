package main

import (
	"context"
	"fmt"
	"time"
)

const (
	Producer = "Producer"
	Consumer = "Consumer"
)

func main() {
	duration := 5 * time.Second // Установка времени, поле которого произойдет отмена контеста
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(duration))
	ch := make(chan int)
	defer cancel()
	i := 0

	go work(ctx, ch, Producer, i)
	go work(ctx, ch, Consumer, i)
	// Ждем конца работы и корректного умирания горутин
	time.Sleep(duration + time.Second)
}

func work(ctx context.Context, ch chan int, worker string, i int) {
	var work string
	if worker == Producer {
		work = "Send %d\t"
	} else {
		work = "Accept %d\n"
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Поймал отмену контекста\n", worker)
			return
		default:
			if worker == Producer {
				fmt.Printf(work, i)
				ch <- i
				i++
			} else {
				fmt.Printf(work, <-ch)
			}
		}
		// Для того чтобы не было слишком быстро
		time.Sleep(200 * time.Millisecond)
	}
}
