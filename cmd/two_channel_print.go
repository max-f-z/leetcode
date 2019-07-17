package main

import (
	"fmt"
	"sync"
)

func twoChannelPrint() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for a := range ch1 {
			if a >= 11 {
				close(ch2)
				close(ch1)
				wg.Done()
			} else {
				fmt.Printf("goroutine 1: %d\n", a)
				a++
				ch2 <- a
			}
		}
	}()

	go func() {
		for b := range ch2 {
			fmt.Printf("goroutine 2: %d\n", b)
			b++
			ch1 <- b
		}
	}()
	ch1 <- 1

	wg.Wait()
}
