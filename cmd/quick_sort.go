package main

// leetcode 912 sort and array
//
//lint:ignore U1000 unused
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	quickSortHelper(&nums, 0, len(nums)-1)
	return nums
}

func quickSortHelper(nums *[]int, start, end int) {
	if start < end {
		pivot := (start + end) / 2

		(*nums)[pivot], (*nums)[start] = (*nums)[start], (*nums)[pivot]

		i, j := start+1, end

		for i <= j {
			if (*nums)[i] <= (*nums)[start] {
				i++
			} else if (*nums)[j] >= (*nums)[start] {
				j--
			} else {
				(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
				i++
				j--
			}
		}

		(*nums)[j], (*nums)[start] = (*nums)[start], (*nums)[j]

		quickSortHelper(nums, start, j-1)
		quickSortHelper(nums, j+1, end)
	}
}
