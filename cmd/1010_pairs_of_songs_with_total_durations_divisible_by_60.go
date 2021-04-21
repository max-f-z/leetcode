package main

func numPairsDivisibleBy60(time []int) int {
	cols := map[int]int{}

	count := 0
	for i := 0; i < len(time); i++ {
		tmp := time[i] % 60
		if _, ok := cols[(60-tmp)%60]; ok {
			count += cols[(60-tmp)%60]
		}
		cols[tmp]++
	}

	return count
}
