package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 只可以讀，不可以寫 (receive-only)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// 只可以寫，不可以讀 (send-only)
	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		wg.Done()
	}(ch)

	wg.Wait()

	/*
		output:
			100
			101
	*/
}
