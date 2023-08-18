package main

func maxSum(nums []int) int {
	ref := map[int]int{}

	max := -1

	for i := 0; i < len(nums); i++ {
		maximum := calculateMax(nums[i])

		if val, ok := ref[maximum]; ok {
			if max < val+nums[i] {
				max = val + nums[i]
			}

			if val > nums[i] {
				continue
			}
		}

		ref[maximum] = nums[i]
	}

	return max
}

func calculateMax(num int) int {
	maximum := 0

	if num == 0 {
		return 0
	}

	for num > 0 {
		tmp := num % 10
		if maximum < tmp {
			maximum = tmp
		}
		num = num / 10
	}

	return maximum
}
