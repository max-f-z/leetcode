package main

//lint:ignore U1000 unused
func multiply(A [][]int, B [][]int) [][]int {
	res := make([][]int, len(A))

	if len(B) == 0 {
		return res
	}
	for i := 0; i < len(A); i++ {
		res[i] = make([]int, len(B[0]))
		for j := 0; j < len(B[0]); j++ {
			sum := 0
			for k := 0; k < len(B); k++ {
				sum += A[i][k] * B[k][j]
			}
			res[i][j] = sum
		}
	}

	return res
}
