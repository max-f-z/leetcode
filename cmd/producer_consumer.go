package main

import (
	"fmt"
	"sync"
	"time"
)

var buf chan int

func demoProducerConsumer() {
	buf = make(chan int, 64)
	go produce(3, buf)
	go produce(5, buf)

	wg := &sync.WaitGroup{}
	wg.Add(100)

	go consume(buf, wg)

	wg.Wait()
}

func produce(factor int, out chan int) {
	for i := 1; ; i++ {
		out <- i * factor
	}
}

func consume(in chan int, wg *sync.WaitGroup) {
	for v := range in {
		fmt.Println(v)
		wg.Done()
		time.Sleep(100 * time.Millisecond)
	}
}
