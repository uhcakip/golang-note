package main

import (
	"fmt"
	"time"
)

func main() {
	jobChan := make(chan int, 1)
	jobCount := 3

	// 啟動 worker
	go worker(jobChan)

	// 將 job 寫入到 channel
	for i := 0; i < jobCount; i++ {
		enqueued := enqueue(i, jobChan)
		fmt.Println(enqueued)
	}

	// 另起 goroutine 將 job 丟到背景等待被寫入
	go func() {
		jobChan <- 3
	}()

	fmt.Println("waiting ...")
	time.Sleep(5 * time.Second)

	/*
		output:
			true
			false
			false
			waiting ...
			job 0 start
			job 0 end
			job 3 start
			job 3 end
	*/
}

func worker(jobChan <-chan int) {
	for i := range jobChan {
		fmt.Println("job", i, "start")
		time.Sleep(2 * time.Second)
		fmt.Println("job", i, "end")
	}
}

func enqueue(job int, jobChan chan<- int) (result bool) {
	select {
	case jobChan <- job: // job 有成功被寫入
		result = true
		return
	default:
		return
	}
}
