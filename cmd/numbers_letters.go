package main

import (
	"fmt"
	"sync"
)

func numsLetters() {
	nums, letters := make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-nums:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letters <- true
				break
			default:
				break
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
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
			default:
			}

		}
	}(wg)

	nums <- true

	wg.Wait()
}
