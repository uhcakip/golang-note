package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Consumer struct {
	// tmpChan store pending jobs
	tmpChan chan int

	// jobChan store jobs ready to process
	jobChan chan int
}

func (c *Consumer) enqueue(job int) {
	select {
	case c.tmpChan <- job:
		fmt.Println("enqueue job", job)
	default:
	}
}

func (c *Consumer) processJob(workerNum, job int) {
	n := getRandInt()
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("worker", workerNum, "cost", n, "seconds to process job", job)
}

func (c *Consumer) startWorker(workerNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start worker", workerNum)

	for job := range c.jobChan {
		c.processJob(workerNum, job)
	}

	fmt.Println("stop worker", workerNum)
}

// 不會等到 job 做完
/*
func (c *Consumer) startWorker(ctx context.Context, workerNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start worker", workerNum)

	for {
		select {
		case job := <-c.jobChan:
			if err := ctx.Err(); err != nil {
				fmt.Printf("context error %v (worker %d job %d)\n", err, workerNum, job)
				return
			}
			c.processJob(workerNum, job)

		case <-ctx.Done():
			fmt.Println("stop worker", workerNum)
			return
		}
	}
}
*/

func (c *Consumer) start(ctx context.Context) {
	for {
		select {
		case job := <-c.tmpChan:
			select {
			case c.jobChan <- job:
				/*
					default:
						fmt.Println("stop process job", job, "due to closed jobChan")
				*/
			}
			if ctx.Err() != nil {
				close(c.jobChan)
				return
			}

		case <-ctx.Done():
			close(c.jobChan)
			return
		}
	}
}

func withContextFunc(parentCtx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
			f()
		}
	}()

	return ctx
}

func getRandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

const poolSize = 5

func main() {
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(poolSize)

	c := &Consumer{
		tmpChan: make(chan int, 10),
		jobChan: make(chan int, poolSize),
	}

	ctx := withContextFunc(context.Background(), func() {
		fmt.Println("got shutdown signal")
		wg.Wait()
		close(done)
	})

	go c.start(ctx)

	for i := 0; i < poolSize; i++ {
		go c.startWorker(i, wg)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			c.enqueue(i)
		}
	}()

	<-done
	fmt.Println("done")
}
