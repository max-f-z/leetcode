package main

func majorityElement(nums []int) int {
	freq := map[int]int{}

	cnt := -1
	ans := -1
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if freq[nums[i]] > cnt {
			cnt = freq[nums[i]]
			ans = nums[i]
		}
	}

	return ans
}
