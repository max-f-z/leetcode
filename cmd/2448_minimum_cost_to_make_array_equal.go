package main

//lint:ignore U1000 unused
func minCostEqual(nums []int, cost []int) int64 {
	var ans int64

	minEl, maxEl := 1000001, 0
	for i := 0; i < len(nums); i++ {
		if minEl > nums[i] {
			minEl = nums[i]
		}
		if maxEl < nums[i] {
			maxEl = nums[i]
		}
	}

	if minEl == maxEl {
		return ans
	}

	l, r := minEl, maxEl
	ans = calCost(nums, cost, nums[0])

	for l < r {
		mid := (l + r) / 2
		cost1 := calCost(nums, cost, mid)
		cost2 := calCost(nums, cost, mid+1)

		ans = cost1
		if ans > cost2 {
			ans = cost2
		}

		if cost1 > cost2 {
			l = mid + 1
		} else {
			r = mid
		}

	}

	return ans
}

func calCost(nums []int, cost []int, target int) int64 {
	var ans int64

	for i := 0; i < len(nums); i++ {
		val := (nums[i] - target) * cost[i]
		if val < 0 {
			val *= -1
		}
		ans += int64(val)
	}

	return ans
}
