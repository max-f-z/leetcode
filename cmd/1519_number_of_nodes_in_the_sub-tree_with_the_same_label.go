package main

func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := map[int]*[]int{}

	for _, v := range edges {
		if graph[v[0]] == nil {
			graph[v[0]] = &[]int{}
		}
		if graph[v[1]] == nil {
			graph[v[1]] = &[]int{}
		}

		(*graph[v[0]]) = append((*graph[v[0]]), v[1])
		(*graph[v[1]]) = append((*graph[v[1]]), v[0])
	}

	visited := make([]int, n)

	ans := make([]int, n)

	countSubTreesHelper(0, graph, &ans, &visited, labels)

	return ans
}

func countSubTreesHelper(cur int, graph map[int]*[]int, ans *[]int, visited *[]int, labels string) map[byte]int {
	(*visited)[cur] = 1
	vals := map[byte]int{}

	if len(*graph[cur]) == 0 {
		(*ans)[cur] = 1
		vals[labels[cur]]++
		return vals
	}

	for _, v := range *graph[cur] {
		if (*visited)[v] == 0 {
			subVals := countSubTreesHelper(v, graph, ans, visited, labels)
			for k, v := range subVals {
				vals[k] += v
			}
		}
	}

	vals[labels[cur]]++
	(*ans)[cur] = vals[labels[cur]]
	return vals
}
