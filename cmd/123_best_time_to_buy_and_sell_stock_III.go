package main

func maxProfit3(prices []int) int {
	profit := 0

	if len(prices) <= 1 {
		return 0
	}

	left := make([]int, len(prices))
	right := make([]int, len(prices))
	left[0] = 0
	right[len(prices)-1] = 0

	l := prices[0]
	r := prices[len(prices)-1]

	for i := 1; i < len(prices); i++ {
		if prices[i] < l {
			l = prices[i]
		}

		if prices[i]-l > left[i-1] {
			left[i] = prices[i] - l
		} else {
			left[i] = left[i-1]
		}

	}

	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > r {
			r = prices[i]
		}

		if r-prices[i] > right[i+1] {
			right[i] = r - prices[i]
		} else {
			right[i] = right[i+1]
		}
	}

	for i := 0; i < len(prices)-1; i++ {
		if left[i]+right[i+1] > profit {
			profit = left[i] + right[i+1]
		}
	}

	if profit < left[len(prices)-1] {
		profit = left[len(prices)-1]
	}

	return profit
}
