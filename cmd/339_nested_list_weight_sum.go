package main

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool {}

func (n NestedInteger) GetInteger() int {}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {}

func depthSum(nestedList []*NestedInteger) int {
	res := 0
	depthSumHelper(nestedList, &res, 1)
	return res
}

func depthSumHelper(num []*NestedInteger, res *int, layer int) {
	for _, v := range num {
		if v.IsInteger() {
			*res = *res + v.GetInteger()*layer
			continue
		} else {
			tmp := 0
			depthSumHelper(v.GetList(), &tmp, layer+1)
			*res += tmp
			continue
		}
	}
}
