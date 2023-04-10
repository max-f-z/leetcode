package main

//lint:ignore U1000 unused
func intersection(nums1 []int, nums2 []int) []int {
	resM := map[int]bool{}

	tmp := map[int]bool{}
	for i := 0; i < len(nums1); i++ {
		tmp[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if tmp[nums2[i]] {
			resM[nums2[i]] = true
		}
	}

	res := []int{}
	for k := range resM {
		res = append(res, k)
	}

	return res
}
