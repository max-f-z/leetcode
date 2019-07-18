package main

func numTrees(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			res[i] += res[j] * res[i-1-j]
		}
	}

	// fmt.Println(res)
	return res[n]
}
