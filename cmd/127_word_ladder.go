package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	l := len(wordList)

	target := -1

	for i := 0; i < l; i++ {
		if endWord == wordList[i] {
			target = i
			break
		}
	}

	if target == -1 {
		return 0
	}

	graph := make([][]int, l)
	for i := 0; i < l; i++ {
		graph[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			if graph[i][j] == 0 {
				if wordLadderHelper(wordList[i], wordList[j]) {
					graph[i][j] = 1
					graph[j][i] = 1
				} else {
					graph[i][j] = -1
					graph[j][i] = -1
				}
			}
		}
	}

	queue := []int{}
	visited := map[int]bool{}
	step := 0

	for i := 0; i < l; i++ {
		if wordList[i] == beginWord {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		for i := 0; i < l; i++ {
			if wordLadderHelper(beginWord, wordList[i]) {
				if wordList[i] == endWord {
					return 2
				}

				queue = append(queue, i)
			}
		}
		step++
	}

	if len(queue) == 0 {
		return 0
	}

	for len(queue) > 0 {
		step++
		newQ := []int{}

		for i := 0; i < len(queue); i++ {
			for j := 0; j < l; j++ {
				if graph[queue[i]][j] == 1 && !visited[j] {
					if wordList[j] == endWord {
						return step + 1
					}
					newQ = append(newQ, j)
					visited[j] = true
				}
			}
		}

		queue = newQ
	}
	return 0
}

func wordLadderHelper(s1, s2 string) bool {
	cnt := 0

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		} else if cnt == 0 {
			cnt++
			continue
		}
		return false
	}

	return true
}
