package main

func checkSubarraySum(nums []int, k int) bool {
	// map key sum % k, value first occurance position
	seen := make(map[int]int)

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum = sum % k

		if sum == 0 && i > 0 {
			return true
		}

		if firstOccurance, ok := seen[sum]; ok {
			if i-firstOccurance >= 2 {
				return true
			}
		} else {
			seen[sum] = i
		}
	}
	return false
}
