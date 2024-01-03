package main

func winnerOfGame(colors string) bool {
	opts := make(map[byte]int, 0)
	opts['A'] = 0
	opts['B'] = 0

	cur := colors[0]
	cnt := 1

	for i := 1; i < len(colors); i++ {
		if cur == colors[i] {
			cnt++
			continue
		} else {
			if cnt >= 3 {
				opts[cur] += cnt - 2
			}

			cur = colors[i]
			cnt = 1
		}
	}

	if cnt >= 3 {
		opts[cur] += cnt - 2
	}

	return opts['A'] > opts['B']
}
