package main

func splitArray(nums []int, m int) int {
	l := 0
	h := 0

	for _, v := range nums {
		h += v
		if v > l {
			l = v
		}
	}

	for l < h {
		mid := (l + h) / 2

		temp := 0
		cnt := 1
		for _, v := range nums {
			temp += v
			if temp > mid {
				temp = v
				cnt++
			}
		}

		if cnt > m {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return l
}
