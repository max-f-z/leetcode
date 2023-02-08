package main

func search33(nums []int, target int) int {
	n := len(nums)

	zeroIdx := -1

	l, r := 0, n-1

	for l < r {
		pivot := (l + r) / 2

		if nums[pivot] > nums[r] {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	zeroIdx = l

	l, r = zeroIdx, zeroIdx+n-1

	for l <= r {
		pivot := (l + r) / 2

		if nums[pivot%n] > target {
			r = pivot - 1
		} else if nums[pivot%n] < target {
			l = pivot + 1
		} else {
			return pivot % n
		}

	}

	return -1
}
