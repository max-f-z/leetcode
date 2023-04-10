package main

type Vector2D struct {
	data []int
	idx  int
}

func Constructor3(v [][]int) Vector2D {
	res := Vector2D{}
	res.idx = -1
	res.data = []int{}
	for _, j := range v {
		res.data = append(res.data, j...)
	}
	return res
}

//lint:ignore ST1006 this
func (this *Vector2D) Next() int {
	this.idx++
	return this.data[this.idx]
}

//lint:ignore ST1006 this
func (this *Vector2D) HasNext() bool {
	return this.idx+1 < len(this.data)
}
