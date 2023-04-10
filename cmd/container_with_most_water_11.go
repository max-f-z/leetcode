package main

//lint:ignore U1000 unused
func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			min := height[j]
			if height[i] < height[j] {
				min = height[i]
			}
			area := (j - i) * min
			if area > max {
				max = area
			}
		}
	}

	return max
}
