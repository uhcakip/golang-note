package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 3; i++ {
		go work(ctx, i)
	}

	time.Sleep(5 * time.Second)
	println("stop goroutines")
	cancel()
}

func work(ctx context.Context, num int) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("stop worker", num)
				return
			default:
				println("worker", num, "still working")
				time.Sleep(time.Second)
			}
		}
	}()
}
