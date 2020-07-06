package main

func numberOfPatterns(m int, n int) int {
	p := 0
	dfs(m, n, 0, -1, 1, make([]bool, 10), &p)
	q := 0
	dfs(m, n, 0, -1, 2, make([]bool, 10), &q)
	r := 0
	dfs(m, n, 0, -1, 5, make([]bool, 10), &r)
	return 4*p + 4*q + r
}

func dfs(m, n, l, parent, current int, visited []bool, result *int) {
	if l == n {
		return
	}
	if visited[current] {
		return
	}

	if parent != -1 {
		sum := parent + current
		if sum%2 == 0 {
			h := sum / 2
			if (h == 2 && ((current == 1 && parent == 3) || (parent == 1 && current == 3))) ||
				(h == 5 && ((current == 4 && parent == 6) || (parent == 4 && current == 6))) ||
				(h == 8 && ((current == 7 && parent == 9) || (parent == 7 && current == 9))) ||
				(h == 4 && ((current == 1 && parent == 7) || (parent == 1 && current == 7))) ||
				(h == 5 && ((current == 2 && parent == 8) || (parent == 2 && current == 8))) ||
				(h == 6 && ((current == 3 && parent == 9) || (parent == 3 && current == 9))) ||
				(h == 5 && ((current == 1 && parent == 9) || (parent == 1 && current == 9))) ||
				(h == 5 && ((current == 7 && parent == 3) || (parent == 7 && current == 3))) {
				if !visited[h] {
					return
				}
			}
		}
	}

	visited[current] = true
	if l+1 >= m {
		*result++
	}
	for i := 1; i <= 9; i++ {
		dfs(m, n, l+1, current, i, visited, result)
	}
	visited[current] = false
}
