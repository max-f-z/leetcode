package main

import "container/list"

type ProductOfNumbers struct {
	els *list.List
}

func Constructor1352() ProductOfNumbers {
	pn := ProductOfNumbers{
		els: list.New(),
	}
	return pn
}

func (this *ProductOfNumbers) Add(num int) {
	this.els.PushFront(num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	val := 1

	el := this.els.Front()

	for i := 0; i < k; i++ {
		tmp := el.Value.(int)
		val = val * tmp
		el = el.Next()
	}

	return val
}
