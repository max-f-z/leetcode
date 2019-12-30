package main

// PredictTheWinner -
func PredictTheWinner(nums []int) bool {

	return optimal(nums, 0, len(nums)-1, true) >= 0
}

func optimal(nums []int, i, j int, turn bool) int {
	if i == j {
		if turn {
			return nums[i]
		} else {
			return nums[i] * -1
		}
	}

	if i > j {
		return 0
	}

	if turn {
		score1 := optimal(nums, i+1, j, false) + nums[i]
		score2 := optimal(nums, i, j-1, false) + nums[j]

		if score1 > score2 {
			return score1
		} else {
			return score2
		}
	} else {
		score1 := optimal(nums, i+1, j, true) - nums[i]
		score2 := optimal(nums, i, j-1, true) - nums[j]

		if score1 > score2 {
			return score2
		} else {
			return score1
		}
	}
}
