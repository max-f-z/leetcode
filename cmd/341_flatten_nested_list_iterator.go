package main

type NestedIterator struct {
	vals []int
	idx  int
}

func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	n := NestedIterator{
		vals: []int{},
		idx:  -1,
	}

	queue := []*NestedInteger{}

	queue = append(queue, nestedList...)

	for {
		newQ := []*NestedInteger{}

		flag := true

		for _, v := range queue {
			if v.IsInteger() {
				newQ = append(newQ, v)
			} else {
				flag = false
				newQ = append(newQ, v.GetList()...)
			}
		}

		if flag {
			break
		}

		queue = newQ
	}

	for _, v := range queue {
		n.vals = append(n.vals, v.GetInteger())
	}

	return &n
}

//lint:ignore ST1006 this
func (this *NestedIterator) Next() int {
	this.idx++
	return this.vals[this.idx]
}

//lint:ignore ST1006 this
func (this *NestedIterator) HasNext() bool {
	return this.idx+1 < len(this.vals)
}
