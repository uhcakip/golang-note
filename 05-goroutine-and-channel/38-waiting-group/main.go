package main

import (
	"fmt"
	"sync"
)

func printNumber(i int, wg *sync.WaitGroup) {
	fmt.Printf("job %d done \n", i)
	// 剩餘 job 數 -1
	wg.Done()
}

func main() {
	fmt.Println("start")
	wg := &sync.WaitGroup{}
	// 有 3 個 job 要執行
	wg.Add(3)

	go printNumber(1, wg)
	go printNumber(2, wg)
	go printNumber(3, wg)

	// 等 job 都做完後 main 才會繼續往下執行
	wg.Wait()
	fmt.Println("end")

	/*
		output:
			start
			job 3 done
			job 2 done
			job 1 done
			end
	*/
}
