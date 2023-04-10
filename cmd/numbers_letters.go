package main

import (
	"fmt"
	"sync"
)

func numsLetters() {
	nums, letters := make(chan bool), make(chan bool)

	go func() {
		i := 1
		//lint:ignore S1000 allow
		for {
			select {
			case <-nums:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letters <- true
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		//lint:ignore S1000 allow
		for {
			select {
			case <-letters:
				fmt.Printf(str[i : i+2])
				i += 2
				if i >= 25 {
					i = 0
					wg.Done()
					break
				}
				nums <- true
			}

		}
	}(wg)

	nums <- true

	wg.Wait()
}
