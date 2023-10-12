package main

type MountainArray struct {
}

//lint:ignore ST1006 this
func (this *MountainArray) get(index int) int {
	return 1
}

//lint:ignore ST1006 this
func (this *MountainArray) length() int {
	return 1
}

//lint:ignore U1000 unused
func findInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 0, mountainArr.length()-1

	for l < r {
		idx := (l + r) / 2
		val := mountainArr.get(idx)
		valNext := mountainArr.get(idx + 1)
		if val < valNext {
			l = idx + 1
		} else {
			r = idx
		}
	}

	peak := l

	l, r = 0, peak-1
	for l < r {
		idx := (l + r) / 2
		val := mountainArr.get(idx)
		if val == target {
			return idx
		} else if val < target {
			l = idx + 1
		} else {
			r = idx
		}
	}

	val := mountainArr.get(l)
	if val == target {
		return l
	}

	l, r = peak, mountainArr.length()
	for l < r {
		idx := (l + r) / 2
		val := mountainArr.get(idx)
		if val == target {
			return idx
		} else if val > target {
			l = idx + 1
		} else {
			r = idx
		}
	}

	return -1
}
