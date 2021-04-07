package main

func maxA(N int) int {
	best := make([]int, N+1)

	for i := 1; i <= N; i++ {
		best[i] = best[i-1] + 1
		for j := 0; j < i-1; j++ {
			best[i] = maxInt(best[i], best[j]*(i-j-1))
		}
	}

	return best[N]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
