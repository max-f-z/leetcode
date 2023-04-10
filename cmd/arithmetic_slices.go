package main

//lint:ignore U1000 unused
func numberOfArithmeticSlices(A []int) int {
	count := 0
	for i := 0; i < len(A)-2; i++ {
		for j := 1; i+j+1 < len(A); j++ {
			if A[i]-A[i+1] == A[i+j]-A[i+j+1] {
				count++
			} else {
				break
			}
		}
	}

	return count
}
