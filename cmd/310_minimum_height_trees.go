package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n < 2 {
		ans := []int{}
		for i := 0; i < n; i++ {
			ans = append(ans, i)
		}
		return ans
	}

	distance := make(map[int]map[int]int)

	for i := 0; i < n; i++ {
		distance[i] = make(map[int]int)
	}

	indegree := map[int]int{}

	for _, edge := range edges {
		distance[edge[0]][edge[1]] = 1
		distance[edge[1]][edge[0]] = 1
		indegree[edge[0]]++
		indegree[edge[1]]++
	}

	queue := []int{}
	seen := map[int]bool{}

	for k, v := range indegree {
		if v == 1 {
			queue = append(queue, k)
			seen[k] = true
		}
	}

	var ans []int

	for len(queue) > 0 {
		newQ := []int{}

		for _, v := range queue {
			for to := range distance[v] {
				if distance[v][to] != 1 {
					continue
				}

				if seen[to] {
					continue
				}

				indegree[to]--

				if indegree[to] != 1 {
					continue
				}

				seen[to] = true
				newQ = append(newQ, to)
			}
		}
		ans = queue
		queue = newQ
	}

	return ans
}
