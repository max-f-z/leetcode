package main

//lint:ignore U1000 unused
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	rowS, rowE, colS, colE := 0, n-1, 0, n-1

	for i := 0; i < n*n; {
		for idx := colS; i < n*n && idx <= colE; idx++ {
			res[rowS][idx] = i + 1
			i++
		}
		rowS++

		for idx := rowS; i < n*n && idx <= rowE; idx++ {
			res[idx][colE] = i + 1
			i++
		}
		colE--

		for idx := colE; i < n*n && idx >= colS; idx-- {
			res[rowE][idx] = i + 1
			i++
		}
		rowE--

		for idx := rowE; i < n*n && idx >= rowS; idx-- {
			res[idx][colS] = i + 1
			i++
		}
		colS++
	}

	return res
}
