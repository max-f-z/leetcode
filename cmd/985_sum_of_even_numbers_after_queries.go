package main

//lint:ignore U1000 unused
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}

	res := []int{}

	for i := 0; i < len(queries); i++ {
		if (A[queries[i][1]]+queries[i][0])%2 == 0 {
			if A[queries[i][1]]%2 == 0 {
				sum += queries[i][0]
			} else {
				sum += A[queries[i][1]] + queries[i][0]
			}
			A[queries[i][1]] = A[queries[i][1]] + queries[i][0]
		} else {
			if A[queries[i][1]]%2 == 0 {
				sum -= A[queries[i][1]]
			}
			A[queries[i][1]] = A[queries[i][1]] + queries[i][0]
		}
		res = append(res, sum)
	}

	return res
}
