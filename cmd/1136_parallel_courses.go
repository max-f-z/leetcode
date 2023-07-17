package main

func minimumSemesters(n int, relations [][]int) int {
	ans := 0

	graph := map[int]int{}

	for i := 1; i <= n; i++ {
		graph[i] = 0
	}

	for _, relation := range relations {
		_, p := relation[0], relation[1]
		graph[p]++
	}

	for len(graph) > 0 {
		nextRemoves := []int{}
		for i := 1; i <= n; i++ {
			if val, ok := graph[i]; ok && val == 0 {
				nextRemoves = append(nextRemoves, i)
			}
		}

		for _, num := range nextRemoves {
			delete(graph, num)
			// graph[num]--
			for _, relation := range relations {
				if relation[0] == num {
					graph[relation[1]]--
				}
			}
		}
		if len(nextRemoves) == 0 && len(graph) > 0 {
			return -1
		}

		ans++
	}

	return ans
}
