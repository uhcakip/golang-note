package main

import (
	"fmt"
	"sync"
)

func main() {
	closeChannel()
	checkIfChannelClosed()
	bufferedChannel()
}

func closeChannel() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(ch <-chan int) {
		// 將 channel 裡面的值透過迴圈讀出來
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		ch <- 102

		// 關閉 channel
		// 此例中不管是 buffered channel 還是 unbuffered channel，若沒有將 channel 關閉一樣會造成 deadlock (讀出 100, 101, 102 後便沒有東西可讀出)
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

	/*
		output:
			100
			101
			102
	*/
}

func checkIfChannelClosed() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, notClosed := <-ch; notClosed {
				fmt.Println(i)
			} else {
				// 要 break 才會在 channel 關閉時跳出程式
				break
			}
		}

		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		// 若往已關閉的 channel 傳值，會噴錯 panic: send on closed channel
		for i := 100; i < 104; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

	/*
		output:
			100
			101
			102
			103
	*/
}

func bufferedChannel() {
	// 設定容量為 10 的 buffered channel
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 讀出 2 個
	go func(ch <-chan int) {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	// 寫入 4 個
	go func(ch chan<- int) {
		ch <- 100
		ch <- 101
		ch <- 102
		ch <- 103
		// buffered channel 在寫入數量沒有超出容量的情況下，即使沒有關閉 channel，值沒有全部讀出來也不會產生 deadlock
		// close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

	/*
		output:
			100
			101
	*/
}
