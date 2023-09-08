package main

//lint:ignore U1000 unused
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row_len := i + 1
		row := make([]int, row_len)

		row[0] = 1
		row[row_len-1] = 1

		idx := 1

		for idx < row_len-1 {
			row[idx] = ans[i-1][idx-1] + ans[i-1][idx]
			idx++
		}

		ans[i] = row
	}

	return ans
}
