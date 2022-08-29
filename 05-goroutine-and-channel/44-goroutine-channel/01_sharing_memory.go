package main

import (
	"fmt"
	"sync"
)

func main() {
	result := sharingMemory(10)
	fmt.Println(len(result), result)
}

func sharingMemory(n int) (ints []int) {
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			ints = append(ints, i)
		}(i)
	}

	wg.Wait()
	return
}
