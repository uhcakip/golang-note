package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	ch := make(chan string, 0)

	defer func() {
		close(ch)
	}()

	go func() {
	LOOP: // 自定義的 label
		for {
			time.Sleep(time.Second)
			fmt.Println(i)
			i++

			select {
			case m := <-ch:
				fmt.Println(m)
				break LOOP
			default: // 沒有 default 的話會 block 在第一個 case 不動
			}
		}
	}()

	time.Sleep(4 * time.Second)
	ch <- "Stop"

	/*
		output:
			0
			1
			2
			3
			Stop
	*/
}
