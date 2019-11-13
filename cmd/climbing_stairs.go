package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a := make([]int, n+1)

	a[1] = 1
	a[2] = 2

	for i := 3; i < n+1; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n]
}
