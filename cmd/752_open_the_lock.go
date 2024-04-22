package main

func openLock(deadends []string, target string) int {
	deadendsMap := map[string]bool{}
	for _, str := range deadends {
		deadendsMap[str] = true
	}

	if deadendsMap["0000"] {
		return -1
	}

	visited := map[string]bool{}
	visited["0000"] = true
	queue := []string{"0000"}

	steps := 0

	for len(queue) > 0 {
		newQ := []string{}
		for _, item := range queue {

			if item == target {
				return steps
			}

			for i := 0; i < len(item); i++ {
				arr := []byte(item)

				arr[i] = calNext(arr[i], 1)

				strP := string(arr)

				arr[i] = calNext(arr[i], -2)
				strN := string(arr)

				if !visited[strP] && !deadendsMap[strP] {
					newQ = append(newQ, strP)
				}

				if !visited[strN] && !deadendsMap[strN] {
					newQ = append(newQ, strN)
				}

				visited[strP] = true
				visited[strN] = true
			}
		}

		steps += 1
		queue = newQ
	}

	return -1
}

func calNext(c byte, offset int) byte {
	c = byte(int(c) + offset)

	if c > '9' {
		c -= 10
	}

	if c < '0' {
		c += 10
	}
	return c
}
