package main

func candy(ratings []int) int {
	n := len(ratings)
	candiesA := make([]int, n)

	candiesA[0] = 1

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candiesA[i] = candiesA[i-1] + 1
		} else {
			candiesA[i] = 1
		}
	}

	candiesB := make([]int, n)
	candiesB[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candiesB[i] = candiesB[i+1] + 1
		} else {
			candiesB[i] = 1
		}
	}

	sum := 0

	for i := 0; i < n; i++ {
		if candiesA[i] > candiesB[i] {
			sum += candiesA[i]
		} else {
			sum += candiesB[i]
		}
	}

	return sum
}
