package main

func mergeSorted(nums1 []int, m int, nums2 []int, n int) {
	idx1 := m - 1
	idx2 := n - 1

	idx3 := m + n - 1

	for idx3 >= 0 {
		if idx1 < 0 {
			nums1[idx3] = nums2[idx2]
			idx2--
			idx3--
			continue
		}
		if idx2 < 0 {
			nums1[idx3] = nums1[idx1]
			idx1--
			idx3--
			continue
		}

		if nums1[idx1] > nums2[idx2] {
			nums1[idx3] = nums1[idx1]
			idx1--
		} else {
			nums1[idx3] = nums2[idx2]
			idx2--
		}
		idx3--
	}
}
