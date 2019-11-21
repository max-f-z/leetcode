package main

import "math"

type MinStack struct {
	content []int
	len     int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	s := MinStack{}
	s.content = make([]int, 0)
	s.len = 0
	return s
}

func (this *MinStack) Push(x int) {
	this.len++
	this.content = append(this.content, x)
}

func (this *MinStack) Pop() {
	this.content = this.content[:len(this.content)-1]
	this.len--
}

func (this *MinStack) Top() int {
	return this.content[len(this.content)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt64
	for i := 0; i < this.len; i++ {
		if this.content[i] < min {
			min = this.content[i]
		}
	}
	return min
}
