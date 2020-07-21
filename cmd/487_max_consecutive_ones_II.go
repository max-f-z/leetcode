package main

func findMaxConsecutiveOnes(nums []int) int {
	start := 0
	end := 0
	numOfZeros := 0
	max := 0

	for end < len(nums) {
		if numOfZeros > 0 && nums[end] == 0 {
			for nums[start] != 0 {
				start++
			}
			start++
			numOfZeros--
		}

		if nums[end] == 0 {
			numOfZeros++
		}
		if end-start+1 > max {
			max = end - start + 1
		}
		end++
	}

	return max
}
