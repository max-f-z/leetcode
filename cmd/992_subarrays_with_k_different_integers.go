package main

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}

// hunch using sliding window
// observation the idx can represent the number of sub arrays that ends at the idx
// p2 - p1 + 1 represent the number of subarrays that ends at p2 and starts after p1 (include p1)
// can represent number of sub arrays at most K
// sub arrays exactly K = sub array at most K - sub array at most K-1
func subarraysWithAtMostKDistinct(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	p1, p2 := 0, 0

	cnt := map[int]int{}
	// distinct := 0

	ans := 0

	for p2 = 0; p2 < len(nums); p2++ {
		cnt[nums[p2]]++

		for len(cnt) > k {
			cnt[nums[p1]]--
			if cnt[nums[p1]] == 0 {
				delete(cnt, nums[p1])
			}
			p1++
		}

		ans += p2 - p1 + 1
	}

	return ans
}
