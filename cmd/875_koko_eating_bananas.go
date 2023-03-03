package main

func minEatingSpeed(piles []int, h int) int {
	maxPerPile := 0

	for i := 0; i < len(piles); i++ {
		if maxPerPile < piles[i] {
			maxPerPile = piles[i]
		}
	}

	l, r := 1, maxPerPile

	for l < r {
		mid := (l + r) / 2
		if eatAll(piles, h, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func eatAll(piles []int, h int, batch int) bool {
	cnt := 0

	for i := 0; i < len(piles); i++ {
		if piles[i]%batch == 0 {
			cnt += piles[i] / batch
		} else {
			cnt += piles[i]/batch + 1
		}
	}

	return cnt <= h
}
