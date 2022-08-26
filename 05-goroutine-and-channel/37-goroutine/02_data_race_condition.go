package main

import (
	"fmt"
	"time"
)

func main() {
	raceCondition()
	solveRaceCondition()
}

func raceCondition() {
	msg := "Let's go"

	go func() {
		fmt.Println(msg)
	}()

	// goroutine 還沒執行，msg 就被換成了 "Let's go go go"
	msg = "Let's go go go"
	time.Sleep(1 * time.Second)

	/*
		output:
			Let's go go go
	*/
}

func solveRaceCondition() {
	msg := "Let's go"

	// 讓 goroutine 接收外面的值
	go func(text string) {
		fmt.Println(text)
	}(msg)

	msg = "Let's go go go"
	time.Sleep(1 * time.Second)

	/*
		output:
			Let's go
	*/
}
