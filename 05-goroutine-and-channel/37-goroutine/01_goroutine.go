package main

import (
	"fmt"
)

func main() {
	// goroutine 會將 printNumber 丟到背景執行
	go printNumber(1)
	go printNumber(2)
	go printNumber(3)

	// goroutine 還來不及執行 main 就已經結束了，要睡一下才會看到輸出結果
	// time.Sleep(1 * time.Second)
}

func printNumber(i int) {
	fmt.Println(i)
}
