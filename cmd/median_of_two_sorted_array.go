package main

//lint:ignore U1000 unused
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := merge(nums1, nums2)

	mid := int(len(nums) / 2)
	if len(nums)%2 == 0 {
		return float64((nums[mid-1] + nums[mid])) / 2
	} else {
		return float64(nums[mid])
	}
}

func merge(nums1 []int, nums2 []int) []int {
	merge := []int{}

	a := 0
	b := 0
	for i := 0; i < len(nums1)+len(nums2); i++ {
		if a >= len(nums1) {
			merge = append(merge, nums2[b])
			b++
			continue
		}

		if b >= len(nums2) {
			merge = append(merge, nums1[a])
			a++
			continue
		}

		if nums1[a] < nums2[b] {
			merge = append(merge, nums1[a])
			a++
		} else {
			merge = append(merge, nums2[b])
			b++
		}
	}

	return merge
}
