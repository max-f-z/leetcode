package main

import "container/heap"

//lint:ignore U1000 unused
func findKthLargest(nums []int, k int) int {
	mh := maxHeap(nums)
	heap.Init(&mh)

	val := 0

	for i := 0; i < k; i++ {
		val = heap.Pop(&mh).(int)
	}

	return val
}

func findKthLargestII(nums []int, k int) int {
	min, max := 10001, -10001
	cnt := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}

		if max < nums[i] {
			max = nums[i]
		}
		cnt[nums[i]]++
	}

	for i := max; i >= min; i-- {
		if cnt[i] == 0 {
			continue
		}

		if k-cnt[i] > 0 {
			k = k - cnt[i]
			continue
		}

		if k-cnt[i] <= 0 {
			return i
		}
	}

	return -1
}
