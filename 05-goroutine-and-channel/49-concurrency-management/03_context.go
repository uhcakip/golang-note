package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, name string) {
	go bar(ctx, name)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[foo]", name, "exit")
			return
		case <-time.After(time.Second):
			fmt.Println("[foo]", name, "working")
		}
	}
}

func bar(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[bar]", name, "exit")
			return
		case <-time.After(2 * time.Second):
			fmt.Println("[bar]", name, "working")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go foo(ctx, "node")
	time.Sleep(5 * time.Second)
	fmt.Println("connection released, notify foo, bar to exit")
	cancel()
}
