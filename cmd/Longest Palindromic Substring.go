package main

//lint:ignore U1000 unused
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	max := 0
	start := 0
	end := 0

	for loop := 0; loop < len(s); loop++ {
		i := 0
		j := loop
		for index := 0; index < len(s)-loop; index++ {

			if j-i == 0 {
				dp[index+i][index+j] = true
			} else if j-i == 1 && s[index+i] == s[index+j] {
				dp[index+i][index+j] = true
			} else if j-i > 1 && s[index+i] == s[index+j] && dp[index+i+1][index+j-1] {
				dp[index+i][index+j] = true
			} else {
				dp[index+i][index+j] = false
			}

			if dp[index+i][index+j] {
				if j-i > max {
					start = index + i
					end = index + j
					max = j - i
				}
			}
			// fmt.Println(index+i, index+j, dp[index+i][index+j])
		}
		j++
		i++
	}

	// fmt.Println(dp)
	// fmt.Println(max, start, end, s[start:end+1])

	return s[start : end+1]
}
