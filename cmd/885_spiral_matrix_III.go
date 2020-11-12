package main

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, R*C)
	for i := 0; i < R*C; i++ {
		res[i] = make([]int, 2)
	}

	matrix := make([][]int, R)
	for i := 0; i < R; i++ {
		matrix[i] = make([]int, C)
	}

	idxR, idxC := r0, c0
	res[0][0], res[0][1] = r0, c0

	step := 1

	for i := 1; i < R*C; {
		for idx := idxC + 1; idx <= c0+step; idx++ {
			if isValid855(idxR, idx, R, C) {
				res[i][0], res[i][1] = idxR, idx
				i++
			}
			idxC++
		}

		for idx := idxR + 1; idx <= r0+step; idx++ {
			if isValid855(idx, idxC, R, C) {
				res[i][0], res[i][1] = idx, idxC
				i++
			}
			idxR++
		}

		for idx := idxC - 1; idx >= c0-step; idx-- {
			if isValid855(idxR, idx, R, C) {
				res[i][0], res[i][1] = idxR, idx
				i++
			}
			idxC--
		}

		for idx := idxR - 1; idx >= r0-step; idx-- {
			if isValid855(idx, idxC, R, C) {
				res[i][0], res[i][1] = idx, idxC
				i++
			}
			idxR--
		}

		step++
	}

	return res
}

func isValid855(r, c, R, C int) bool {
	if r >= 0 && r < R && c >= 0 && c < C {
		return true
	}
	return false
}
