package main

import (
	"container/list"
)

type HitCounter struct {
	stack *list.List
}

func Constructor362() HitCounter {
	hc := HitCounter{}
	hc.stack = list.New().Init()
	return hc
}

//lint:ignore ST1006 this
func (this *HitCounter) Hit(timestamp int) {
	this.stack.PushBack(timestamp)
}

//lint:ignore ST1006 this
func (this *HitCounter) GetHits(timestamp int) int {
	for {
		if this.stack.Len() == 0 {
			break
		}
		val := this.stack.Front().Value.(int)
		if timestamp-300 >= val {
			this.stack.Remove(this.stack.Front())
		} else {
			break
		}
	}

	return this.stack.Len()
}
