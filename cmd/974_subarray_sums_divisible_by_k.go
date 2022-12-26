package main

func subarraysDivByK(nums []int, k int) int {
	prefixMod := make([]int, len(nums))
	modGroup := make([]int, k)

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixMod[i] = (sum%k + k) % k
		modGroup[prefixMod[i]]++
	}

	result := 0
	for i := 0; i < k; i++ {
		if modGroup[i] >= 2 {
			result += (modGroup[i] * (modGroup[i] - 1) / 2)
		}
	}

	result += modGroup[0]

	return result
}
