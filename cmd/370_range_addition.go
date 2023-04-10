package main

//lint:ignore U1000 unused
func getModifiedArray(length int, updates [][]int) []int {
	res := make([]int, length)

	for _, v := range updates {
		start := v[0]
		end := v[1]
		update := v[2]

		res[start] += update
		if end < length-1 {
			res[end+1] -= update
		}
	}

	for i := 1; i < length; i++ {
		res[i] += res[i-1]
	}

	return res
}
