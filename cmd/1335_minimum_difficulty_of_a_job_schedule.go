package main

import "math"

// ref: https://leetcode.jp/leetcode-1335-minimum-difficulty-of-a-job-schedule-%E8%A7%A3%E9%A2%98%E6%80%9D%E8%B7%AF%E5%88%86%E6%9E%90/

var memo1335 [][]int

func minDifficulty(jobDifficulty []int, d int) int {
	if d > len(jobDifficulty) {
		return -1
	}

	memo1335 = make([][]int, d+1)
	for i := 0; i < d+1; i++ {
		memo1335[i] = make([]int, len(jobDifficulty))
	}

	minDifficultyHelper(jobDifficulty, 0, d)

	return memo1335[d][0]
}

func minDifficultyHelper(jobDifficulty []int, startIndex int, d int) int {

	if memo1335[d][startIndex] > 0 {
		return memo1335[d][startIndex]
	}

	maxDifficulty := 0
	res := math.MaxInt64

	end := len(jobDifficulty) - d

	for i := startIndex; i <= end; i++ {
		if jobDifficulty[i] > maxDifficulty {
			maxDifficulty = jobDifficulty[i]
		}

		if d > 1 {
			local := maxDifficulty + minDifficultyHelper(jobDifficulty, i+1, d-1)
			if local < res {
				res = local
			}
		}
	}

	if d == 1 {
		res = maxDifficulty
	}

	memo1335[d][startIndex] = res
	return res
}
