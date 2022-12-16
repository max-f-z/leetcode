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

func (this *MyQueue) Pop() int {
	val := (this.s1.Back().Value).(int)
	this.s1.Remove(this.s1.Back())
	return val
}

func (this *MyQueue) Peek() int {
	return this.s1.Back().Value.(int)
}

func (this *MyQueue) Empty() bool {
	if this.s1.Len() == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
