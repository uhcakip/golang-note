package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan bool)

	go func() {
		for {
			select {
			case <-exit:
				fmt.Println("exit")
				return
			case <-time.After(time.Second):
				fmt.Println("monitoring")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	exit <- true
}
