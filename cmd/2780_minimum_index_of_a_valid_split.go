package main

func minimumIndex(nums []int) int {
	first_dom := -1

	n := len(nums)
	first_len := 0
	second_len := n

	first := map[int]int{}
	second := map[int]int{}

	for i := 0; i < len(nums); i++ {
		second[nums[i]] += 1
	}

	for i := 0; i < len(nums); i++ {
		first[nums[i]] += 1
		first_len += 1

		second[nums[i]] -= 1
		second_len -= 1

		if first[nums[i]] > first_len/2 {
			first_dom = nums[i]
		}

		if first[first_dom] <= first_len/2 {
			first_dom = -1
			continue
		}

		if second[first_dom] > second_len/2 {
			return i
		}
	}
	return -1
}
