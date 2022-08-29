package main

import (
	"fmt"
)

func main() {
	result := solveSharingMemoryByChannel(10)
	fmt.Println(len(result), result)
}

func solveSharingMemoryByChannel(n int) (ints []int) {
	ch := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(ch chan<- int, i int) { // 確保 goroutine 只能寫入，不能讀出
			ch <- i
		}(ch, i)
	}

	for i := range ch {
		ints = append(ints, i)
		if len(ints) == n {
			break
		}
	}

	close(ch)
	return
}
