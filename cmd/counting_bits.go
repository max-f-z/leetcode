package main

//lint:ignore U1000 unused
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0

	for i := 1; i < num+1; i++ {
		if i&1 == 0 {
			res[i] = res[i>>1]
		} else {
			res[i] = res[i-1] + 1
		}
	}

	return res
}
