package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	ans := make([]int, n)
	ans[0] = 0
	for i := 1; i < n; i++ {
		ans[i] = -1
	}

	red := map[int][]int{}
	visitedRed := make([]int, n)

	for _, v := range redEdges {
		if red[v[0]] == nil {
			red[v[0]] = []int{}
		}
		red[v[0]] = append(red[v[0]], v[1])
	}

	blue := map[int][]int{}
	visitedBlue := make([]int, n)

	for _, v := range blueEdges {
		if blue[v[0]] == nil {
			blue[v[0]] = []int{}
		}
		blue[v[0]] = append(blue[v[0]], v[1])
	}

	queue := [][]int{{0, 0}, {0, 1}}

	step := 0

	for len(queue) > 0 {
		newQ := [][]int{}

		for _, v := range queue {
			node := v[0]
			color := v[1]

			if color == 0 { // red
				if step < ans[node] || ans[node] == -1 {
					ans[node] = step
					visitedRed[node] = 1
				}
				for _, vv := range red[node] {
					if visitedBlue[vv] != 1 {
						visitedBlue[vv] = 1
						newQ = append(newQ, []int{vv, 1})
					}
				}
			} else { // blue
				if step < ans[node] || ans[node] == -1 {
					ans[node] = step
					visitedBlue[node] = 1
				}
				for _, vv := range blue[node] {
					if visitedRed[vv] != 1 {
						visitedRed[vv] = 1
						newQ = append(newQ, []int{vv, 0})
					}
				}
			}
		}
		step++
		queue = newQ
	}

	return ans
}
