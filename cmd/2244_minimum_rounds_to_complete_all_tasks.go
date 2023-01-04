package main

func minimumRounds(tasks []int) int {
	cnt := map[int]int{}

	for i := 0; i < len(tasks); i++ {
		cnt[tasks[i]]++
	}

	ans := 0
	for _, v := range cnt {
		if v == 1 {
			return -1
		}

		if v%3 == 1 || v%3 == 2 {
			ans += v/3 + 1
			continue
		}

		if v%3 == 0 {
			ans += v / 3
		}
	}

	return ans
}
