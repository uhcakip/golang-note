package main

import (
	"fmt"
	"sync"
)

func main() {
	sample1()
	sample2()
	sample3()
}

func sample1() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			// unbuffered channel 會等到有值寫入 channel 內
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			i := 100
			ch <- i
			// ch <- i 是 pass by value，所以 channel 裡的值不會被改變
			i = 101
			wg.Done()
		}()
	}

	wg.Wait()

	/*
		output:
			100
			100
			100
			100
			100
	*/
}

func sample2() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 101
		wg.Done()
	}()

	go func() {
		i := 100
		ch <- i
		// 將 ch <- 101 讀出
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}()

	wg.Wait()

	/*
		output:
			100
			101
	*/
}

func sample3() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(ch chan int) {
		i := <-ch
		fmt.Println(i)
		ch <- 101
		wg.Done()
	}(ch)

	go func(ch chan int) {
		i := 100
		ch <- i
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	wg.Wait()

	/*
		output:
			100
			101
	*/
}
