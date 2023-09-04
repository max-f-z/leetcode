package main

func countSymmetricIntegers(low int, high int) int {
	cnt := 0

	for i := low; i <= high; i++ {
		if isNumSymmetric(i) {
			cnt++
		}
	}

	return cnt
}

func isNumSymmetric(num int) bool {
	digits := 0
	preSum := []int{}
	sum := 0

	for num > 0 {
		digits++
		sum += num % 10
		preSum = append(preSum, sum)

		num = num / 10

	}

	if digits%2 == 1 {
		return false
	}

	return preSum[digits/2-1] == preSum[digits-1]-preSum[digits/2-1]
}
