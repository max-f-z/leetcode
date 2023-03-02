package main

import "strconv"

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}

	count, idx := 1, 0

	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[idx] {
			count++
		} else {
			if count == 1 {
				idx = idx + 1
			} else {
				nums := []byte(strconv.Itoa(count))
				for j := 0; j < len(nums); j++ {
					chars[idx+1+j] = nums[j]
				}
				idx = idx + len(nums) + 1
			}
			chars[idx] = chars[i]
			count = 1
		}
	}

	if count > 1 {
		nums := []byte(strconv.Itoa(count))
		for j := 0; j < len(nums); j++ {
			chars[idx+1+j] = nums[j]
		}
		idx = idx + len(nums) + 1
	} else {

		idx++
	}

	return idx
}
