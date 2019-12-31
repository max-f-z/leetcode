package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left := make([]int, len(height)-1)
	right := make([]int, len(height)-1)

	water := 0
	max := height[0]

	for i := 1; i < len(height)-1; i++ {
		if height[i] < max {
			left[i] = max - height[i]
		} else {
			max = height[i]
			left[i] = 0
		}
	}

	max = height[len(height)-1]

	for i := len(height) - 2; i >= 0; i-- {
		if height[i] < max {
			right[i] = max - height[i]
		} else {
			max = height[i]
			right[i] = 0
		}
	}

	for i := 0; i < len(height)-1; i++ {
		if left[i] < right[i] {
			water += left[i]
		} else {
			water += right[i]
		}
	}

	return water
}
