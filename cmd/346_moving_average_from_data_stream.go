package main

type MovingAverage struct {
	data []int
	idx  int
	size int
	sum  float64
}

/** Initialize your data structure here. */
func Constructor346(size int) MovingAverage {
	ma := MovingAverage{}
	ma.data = make([]int, size)
	ma.idx = 0
	ma.size = size
	return ma
}

func (this *MovingAverage) Next(val int) float64 {
	this.idx++
	tmp := 0
	this.data[(this.idx-1)%this.size], tmp = val, this.data[(this.idx-1)%this.size]
	this.sum = this.sum - float64(tmp) + float64(val)

	if this.idx <= this.size {
		return this.sum / float64(this.idx)
	}
	return this.sum / float64(this.size)
}
