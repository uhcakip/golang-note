package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	delta := 3

	var wg sync.WaitGroup
	wg.Add(delta)

	for i := 0; i < delta; i++ {
		go func(i int, count int) {
			defer wg.Done()
			fmt.Println("goroutine", i)
			count++
		}(i, count)
	}

	wg.Wait()
	fmt.Println("count", count)
}
