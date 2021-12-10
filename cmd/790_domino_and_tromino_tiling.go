package main

func numTilings(n int) int {
	if n == 1 {
		return 1
	}

	fullCover := make([]int, n+1)
	fullCover[1] = 1
	fullCover[2] = 2

	partialCover := make([]int, n+1)
	partialCover[2] = 1

	for i := 3; i <= n; i++ {
		fullCover[i] = (fullCover[i-1] + fullCover[i-2] + 2*partialCover[i-1]) % 1000000007
		partialCover[i] = (partialCover[i-1] + fullCover[i-2]) % 1000000007
	}

	return fullCover[n]
}
