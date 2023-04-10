package main

//lint:ignore U1000 unused
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m := len(matrix)
	n := len(matrix[0])

	res := make([]int, m*n)

	rowS, rowE, colS, colE := 0, m-1, 0, n-1

	for i := 0; i < m*n; {
		for idx := colS; i < len(res) && idx <= colE; idx++ {
			res[i] = matrix[rowS][idx]
			i++
		}
		rowS++

		for idx := rowS; i < len(res) && idx <= rowE; idx++ {
			res[i] = matrix[idx][colE]
			i++
		}
		colE--

		for idx := colE; i < len(res) && idx >= colS; idx-- {
			res[i] = matrix[rowE][idx]
			i++
		}
		rowE--

		for idx := rowE; i < len(res) && idx >= rowS; idx-- {
			res[i] = matrix[idx][colS]
			i++
		}
		colS++

	}

	return res
}
