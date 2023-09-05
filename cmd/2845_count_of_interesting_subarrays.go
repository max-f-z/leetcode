package main

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n, preSum := len(nums), 0
	var count int64
	hm := make(map[int]int64)
	for i := 0; i <= n; i++ {
		if hm[(modulo+preSum%modulo-k)%modulo] > 0 {
			count += int64(hm[(modulo+preSum%modulo-k)%modulo])
		}
		if hm[preSum%modulo] == 0 {
			hm[preSum%modulo] = 1
		} else {
			hm[preSum%modulo]++
		}
		if i < n && nums[i]%modulo == k {
			preSum += 1
		}
	}
	return count
}
