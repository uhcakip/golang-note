package main

import "fmt"

func main() {
	c := make(chan bool)

	go func() {
		fmt.Println("go go go")
		// 讀出 (此處會等待值寫入)
		// 註解此行會噴錯 fatal error: all goroutines are asleep - deadlock (有寫入但未讀出)
		<-c
	}()

	// 寫入
	c <- true

	/*
		output:
			go go go
	*/
}
