package main

type ZigzagIterator struct {
	v1   []int
	v2   []int
	flag bool
}

func Constructor281(v1, v2 []int) *ZigzagIterator {
	z := &ZigzagIterator{v1: v1, v2: v2, flag: true}
	return z
}

//lint:ignore ST1006 this
func (this *ZigzagIterator) next() int {
	var v int
	if this.flag {
		if len(this.v1) > 0 {
			v, this.v1 = this.v1[0], this.v1[1:]
			this.flag = !this.flag
		} else {
			v, this.v2 = this.v2[0], this.v2[1:]
		}
	} else {
		if len(this.v2) > 0 {
			v, this.v2 = this.v2[0], this.v2[1:]
			this.flag = !this.flag
		} else {
			v, this.v1 = this.v1[0], this.v1[1:]
		}
	}

	return v
}

//lint:ignore ST1006 this
func (this *ZigzagIterator) hasNext() bool {
	if len(this.v1) == 0 && len(this.v2) == 0 {
		return false
	}
	return true
}
