package main

import "container/list"

type MyQueue struct {
	s1 *list.List
	s2 *list.List
}

func Constructor232() MyQueue {
	return MyQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

//lint:ignore ST1006 this
func (this *MyQueue) Push(x int) {
	if this.Empty() {
		this.s1.PushBack(x)
	} else {
		for this.s1.Len() != 0 {
			val := (this.s1.Back().Value).(int)
			this.s2.PushBack(val)
			this.s1.Remove(this.s1.Back())
		}
		this.s2.PushBack(x)

		for this.s2.Len() != 0 {
			val := (this.s2.Back().Value).(int)
			this.s1.PushBack(val)
			this.s2.Remove(this.s2.Back())
		}
	}
}

//lint:ignore ST1006 this
func (this *MyQueue) Pop() int {
	val := (this.s1.Back().Value).(int)
	this.s1.Remove(this.s1.Back())
	return val
}

//lint:ignore ST1006 this
func (this *MyQueue) Peek() int {
	return this.s1.Back().Value.(int)
}

//lint:ignore ST1006 this
func (this *MyQueue) Empty() bool {
	return this.s1.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
