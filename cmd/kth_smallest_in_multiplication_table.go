package main

func findKthNumber(m int, n int, k int) int {
	l := 1
	h := m * n
	for l < h {
		mid := (l + h) / 2
		if !enough(mid, m, n, k) {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return l
}

func enough(x, m, n, k int) bool {
	count := 0
	for i := 1; i <= m; i++ {
		min := x / i
		if min < n {
			count += min
		} else {
			count += n
		}
	}

	return count >= k
}
