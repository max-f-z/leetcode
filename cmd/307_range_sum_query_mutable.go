package main

//https://www.youtube.com/watch?v=8XSfWhZFSBs

type SegmentTreeNode struct {
	s     int
	e     int
	val   int
	left  *SegmentTreeNode
	right *SegmentTreeNode
}

type NumArray struct {
	root *SegmentTreeNode
}

func buildSegTree(start, end int, nums []int) *SegmentTreeNode {
	if start > end {
		return nil
	}
	n := &SegmentTreeNode{}
	n.s = start
	n.e = end
	if start == end {
		n.val = nums[start]
		return n
	}

	mid := start + (end-start)/2
	n.left = buildSegTree(start, mid, nums)
	n.right = buildSegTree(mid+1, end, nums)
	n.val = n.left.val + n.right.val
	return n
}

func Constructor307(nums []int) NumArray {
	n := NumArray{}
	n.root = buildSegTree(0, len(nums)-1, nums)
	return n
}

//lint:ignore ST1006 this
func (this *NumArray) Update(i int, val int) {
	updateHelper(i, val, this.root)
}

func updateHelper(i int, val int, node *SegmentTreeNode) int {
	if node.s == i && node.e == i {
		tmp := val - node.val
		node.val = val
		return tmp
	}

	mid := (node.e-node.s)/2 + node.s
	if i > mid {
		tmp := updateHelper(i, val, node.right)
		node.val += tmp
		return tmp
	}
	tmp := updateHelper(i, val, node.left)
	node.val += tmp

	return tmp
}

//lint:ignore ST1006 this
func (this *NumArray) SumRange(i int, j int) int {
	return sumRangeHelper(i, j, this.root)
}

func sumRangeHelper(s, e int, node *SegmentTreeNode) int {
	if s > e {
		return 0
	}

	if s == node.s && e == node.e {
		return node.val
	}

	mid := (node.s + node.e) / 2

	if s > mid {
		return sumRangeHelper(s, e, node.right)
	} else if e < mid {
		return sumRangeHelper(s, e, node.left)
	} else {
		return sumRangeHelper(s, mid, node.left) + sumRangeHelper(mid+1, e, node.right)
	}
}
