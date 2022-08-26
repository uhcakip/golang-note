package main

import "fmt"

func main() {
	bufferedChannel()
	bufferedChannelDeadlock()
}

func bufferedChannel() {
	// 建立一個容量為 1 的 buffered channel
	c := make(chan bool, 1)

	go func() {
		fmt.Println("go go go")
		// 主程式會繼續往下執行，不需等 channel 的值被寫入或讀出
		// 註解此行不會噴錯
		<-c
	}()

	c <- true

	/*
		output:
			無
	*/
}

func bufferedChannelDeadlock() {
	c := make(chan bool, 1)

	go func() {
		fmt.Println("go go go")
		// <-c
	}()

	// 寫入的數量超出限制容量時，在其他 goroutine (或 main) 沒有讀出的情況下會造成 deadlock
	c <- true
	c <- true

	/*
		output:
			go go go
			fatal error: all goroutines are asleep - deadlock
	*/
}
