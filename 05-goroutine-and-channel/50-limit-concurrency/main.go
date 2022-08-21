package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 限制 worker 數量
	const workerCount = 10
	const jobCount = 100

	var wg sync.WaitGroup
	wg.Add(jobCount)

	jobQueue := make(chan int)
	jobDone := make(chan int)

	// 將 job 塞入 queue
	go func(queue chan<- int) {
		for i := 0; i < jobCount; i++ {
			jobQueue <- i
		}
		close(jobQueue)
	}(jobQueue)

	// 啟動 10 個 worker 處理 job
	for i := 0; i < workerCount; i++ {
		go func(queue <-chan int) {
			for job := range jobQueue {
				defer wg.Done()
				waitTime := rand.Int31n(1000)
				fmt.Println("process job", job, ", wait", waitTime, "ms")
				time.Sleep(time.Duration(waitTime) * time.Millisecond)
				jobDone <- job
			}
		}(jobQueue)
	}

	// 在背景等待，避免 main 被卡住
	go func() {
		wg.Wait()
		close(jobDone)
	}()

	var results []int
	for f := range jobDone {
		fmt.Println("processed job", f)
		results = append(results, f)
	}

	fmt.Println("results", results)
}
