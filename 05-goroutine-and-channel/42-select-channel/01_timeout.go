package main

import (
	"fmt"
	"time"
)

func main() {
	sample1()
	sample2()
}

func sample1() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	ch := make(chan int)

	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout")
	}
}

func sample2() {
	ch := make(chan int)

	select {
	case <-ch:
	case <-time.After(time.Second): // 回傳一個 channel
		println("timeout")

	}
}
