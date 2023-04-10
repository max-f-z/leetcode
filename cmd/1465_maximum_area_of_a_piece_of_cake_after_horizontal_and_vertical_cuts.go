package main

import "sort"

//lint:ignore U1000 unused
func maxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxH := horizontalCuts[0]
	maxV := verticalCuts[0]
	for i := 1; i < len(horizontalCuts); i++ {
		if maxH < horizontalCuts[i]-horizontalCuts[i-1] {
			maxH = horizontalCuts[i] - horizontalCuts[i-1]
		}
	}

	if maxH < h-horizontalCuts[len(horizontalCuts)-1] {
		maxH = h - horizontalCuts[len(horizontalCuts)-1]
	}

	for i := 1; i < len(verticalCuts); i++ {
		if maxV < verticalCuts[i]-verticalCuts[i-1] {
			maxV = verticalCuts[i] - verticalCuts[i-1]
		}
	}

	if maxV < w-verticalCuts[len(verticalCuts)-1] {
		maxV = w - verticalCuts[len(verticalCuts)-1]
	}

	return maxH * maxV % 1000000007
}
