package main

func addToArrayForm(num []int, k int) []int {
	reverseInts(&num)
	op := convertInt(k)

	m, n := len(num), len(op)

	carry := 0

	ans := []int{}

	for p1, p2 := 0, 0; p1 < m || p2 < n; p1, p2 = p1+1, p2+1 {
		sum := 0
		sum += carry
		if p1 < m {
			sum += num[p1]
		}
		if p2 < n {
			sum += op[p2]
		}

		if sum >= 10 {
			carry = 1
			ans = append(ans, sum-10)
		} else {
			ans = append(ans, sum)
			carry = 0
		}
	}

	if carry == 1 {
		ans = append(ans, 1)
	}

	reverseInts(&ans)
	return ans
}

func reverseInts(num *[]int) {
	if len(*num) < 2 {
		return
	}

	for i := 0; i < len(*num)/2; i++ {
		(*num)[i], (*num)[len(*num)-1-i] = (*num)[len(*num)-1-i], (*num)[i]
	}
}

func convertInt(num int) []int {
	ans := []int{}

	for num > 10 {
		ans = append(ans, num%10)
		num = (num - num%10) / 10
	}

	if num != 0 {
		ans = append(ans, num)
	}

	return ans
}
