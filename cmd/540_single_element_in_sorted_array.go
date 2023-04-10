package main

//lint:ignore U1000 unused
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if nums[mid] == nums[mid-1] {
			if mid-1 == 0 || (r-mid)%2 == 1 {
				l = mid + 1
				continue
			} else {
				r = mid - 2
			}
		} else {
			if mid == 0 || (r-mid-1)%2 == 1 {
				l = mid + 2
				continue
			} else {
				r = mid - 1
			}
		}
	}

	return nums[l]
}
