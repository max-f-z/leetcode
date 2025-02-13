package main

import "container/heap"

func maximumSum(nums []int) int {
	sumRef := map[int]*maxHeap{}

	for _, v := range nums {
		ds := digitSum(v)

		if mh, ok := sumRef[ds]; ok {
			heap.Push(mh, v)
		} else {
			mh := &maxHeap{}
			heap.Init(mh)
			heap.Push(mh, v)
			sumRef[ds] = mh
		}
	}

	ans := -1

	for _, mh := range sumRef {
		if mh.Len() >= 2 {
			op1 := heap.Pop(mh).(int)
			op2 := heap.Pop(mh).(int)

			if op1+op2 > ans {
				ans = op1 + op2
			}
		}
	}

	return ans
}

func digitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = (num - num%10) / 10
	}

	return sum
}
