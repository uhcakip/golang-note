package main

import (
	"fmt"
	"sync"
)

func main() {
	result := solveSharingMemoryByLock(10)
	fmt.Println(len(result), result)
}

func solveSharingMemoryByLock(n int) (ints []int) {
	var (
		wg sync.WaitGroup // 確保全部 goroutine 都拿到資料後才結束 func
		mx sync.Mutex
	)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mx.Lock()
			ints = append(ints, i)
			mx.Unlock()
		}(i)
	}

	wg.Wait()
	return
}
