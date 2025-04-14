package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)

	abs := func(num int) int {
		if num < 0 {
			return num * -1
		}

		return num
	}

	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}

	return ans
}
