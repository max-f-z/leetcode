package main

//lint:ignore U1000 unused
func wiggleSort(nums []int) {
	inc := true

	for i := 0; i < len(nums)-1; i++ {
		if inc {
			inc = !inc
			if nums[i] <= nums[i+1] {
				continue
			} else {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		} else {
			inc = !inc
			if nums[i] >= nums[i+1] {
				continue
			} else {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}

	}
}
