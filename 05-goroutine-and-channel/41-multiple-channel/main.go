package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 20
	outChan := make(chan string, 100)
	finishChan := make(chan struct{}) // struct channel 不佔任何記憶體
	errChan := make(chan error, 100)

	wg := &sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(i int, wg *sync.WaitGroup, outChan chan<- string, errChan chan<- error) {
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			outChan <- fmt.Sprintf("Processed job id: %d", i)

			/*
				if i == 15 {
					errChan <- errors.New("error in job 15, break")
				}
			*/

			wg.Done()
		}(i, wg, outChan, errChan)
	}

	go func() {
		wg.Wait()
		close(finishChan)
	}()

Loop:
	for { // 若沒有 break 會一直等待
		select {
		case out := <-outChan:
			fmt.Println("Value from outChan: ", out)
		case <-finishChan:
			fmt.Println("Jobs are all finished, break")
			break Loop
		case err := <-errChan:
			fmt.Println(err.Error())
			break Loop
		case <-time.After(100 * time.Millisecond): // set timeout
			fmt.Println("Timeout, break")
			break Loop
		}
	}
}
