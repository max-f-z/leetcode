package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1

	for l < r {
		val := (l + r) / 2
		if arr[val] < arr[val+1] {
			l = val + 1
		} else {
			r = val
		}
	}

	return l
}
