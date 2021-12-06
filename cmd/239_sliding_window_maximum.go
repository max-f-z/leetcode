package main

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}

	deque := list.New()

	for i := 0; i < k; i++ {
		for deque.Len() != 0 && nums[i] > nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)
	}

	for i := k; i < len(nums); i++ {
		res = append(res, nums[deque.Front().Value.(int)])

		front := deque.Front().Value.(int)
		if i-k == front {
			deque.Remove(deque.Front())
		}

		for deque.Len() != 0 && nums[i] > nums[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
	}

	res = append(res, nums[deque.Front().Value.(int)])
	return res
}
