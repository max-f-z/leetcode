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
		for _, k := range j {
			res.data = append(res.data, k)
		}
	}
	return res
}

func (this *Vector2D) Next() int {
	this.idx++
	return this.data[this.idx]
}

func (this *Vector2D) HasNext() bool {
	return this.idx+1 < len(this.data)
}
