package main

//lint:ignore U1000 unused
func numberOfBeams(bank []string) int {
	ans := 0
	prev := 0
	for _, row := range bank {
		cnt := 0
		for idx := range row {
			if row[idx] == '1' {
				cnt++
			}
		}
		if cnt == 0 {
			continue
		}
		ans += prev * cnt
		prev = cnt
	}

	return ans
}
