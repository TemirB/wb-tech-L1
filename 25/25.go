package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	if d <= 0 {
		return
	}
	<-time.After(d)
}

func main() {
	fmt.Println("start:", time.Now().Format(time.RFC3339Nano))
	sleep(500 * time.Millisecond)
	fmt.Println("done :", time.Now().Format(time.RFC3339Nano))
}
