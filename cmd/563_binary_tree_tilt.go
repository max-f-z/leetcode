package main

import "math"

func findTilt(root *TreeNode) int {
	tilt, _ := findTiltHelper(root)
	return tilt
}

func findTiltHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	var lsum, ltilt, rsum, rtilt, tilt, sum int

	ltilt, lsum = findTiltHelper(root.Left)

	rtilt, rsum = findTiltHelper(root.Right)

	tilt = int(math.Abs(float64(lsum)-float64(rsum))) + ltilt + rtilt

	sum = lsum + rsum + root.Val

	return tilt, sum
}
