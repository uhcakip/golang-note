package main

import (
	"fmt"
	"sync"
)

func main() {
	case1()
	case2()
}

func case1() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func(ch <-chan int) {
		j := <-ch
		fmt.Println(j)
		wg.Done()
	}(ch)

	// 在第二次迴圈時沒有 channel 去讀出 j
	// 在 job 沒有做完的情況下造成了 deadlock (wg.Wait)
	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func(ch chan<- int, j int) {
			ch <- j
			wg.Done()
		}(ch, j)
	}

	wg.Wait()

	/*
		output:
			0
			fatal error: all goroutines are asleep - deadlock
	*/
}

func case2() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		// 沒有 channel 將 `ch <- 102` 讀出來，同上例原因造成 deadlock
		ch <- 102
		wg.Done()
	}(ch)

	wg.Wait()

	/*
		output:
			100
			101
			fatal error: all goroutines are asleep - deadlock! ...
	*/
}
