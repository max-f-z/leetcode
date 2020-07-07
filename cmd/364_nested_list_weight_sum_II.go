package main

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool {
	return true
}

func (n NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func depthSumInverse(nestedList []*NestedInteger) int {
	res := 0
	max := getDepth(nestedList)
	depthSumInverseHelper(nestedList, &res, max)
	return res
}

func depthSumInverseHelper(num []*NestedInteger, res *int, layer int) {
	for _, v := range num {
		if v.IsInteger() {
			*res = *res + v.GetInteger()*layer
			continue
		} else {
			tmp := 0
			depthSumInverseHelper(v.GetList(), &tmp, layer-1)
			*res += tmp
			continue
		}
	}
}

func getDepth(nestedList []*NestedInteger) int {
	max := 1
	for _, v := range nestedList {
		if v.IsInteger() {
			continue
		} else {
			tmp := getDepth(v.GetList()) + 1
			if max < tmp {
				max = tmp
			}
		}
	}
	return max
}
