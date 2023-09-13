package main

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	dx := fx - sx
	dy := fy - sy
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}

	minStep := dx
	if dy > minStep {
		minStep = dy
	}

	if minStep == 0 {
		return t != 1
	}

	return t >= minStep
}
