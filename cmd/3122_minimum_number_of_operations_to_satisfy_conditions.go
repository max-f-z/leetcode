package main

func minimumOperations3122(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	freq := map[int]map[int]int{}

	for i := 0; i < cols; i++ {
		if freq[i] == nil {
			freq[i] = map[int]int{}
		}

		for j := 0; j < rows; j++ {
			freq[i][grid[j][i]] += 1
		}
	}

	mem := map[int]map[int]int{}
	for i := 0; i < cols; i++ {
		mem[i] = map[int]int{}
	}

	for i := 0; i < 10; i++ {
		mem[cols-1][i] = rows - freq[cols-1][i]
	}

	ans := 1000001

	for i := cols - 2; i >= 0; i-- {
		for j := 0; j < 10; j++ {
			current := rows - freq[i][j]
			min := 1000001
			for k, v := range mem[i+1] {
				if k == j {
					continue
				}
				if min > v {
					min = v
				}
			}

			current += min

			mem[i][j] = current
		}
	}

	for i := 0; i < 10; i++ {
		if mem[0][i] < ans {
			ans = mem[0][i]
		}
	}

	return ans
}
