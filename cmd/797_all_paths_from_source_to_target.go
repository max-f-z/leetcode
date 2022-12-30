package main

var ans797 [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	ans797 = [][]int{}

	dfs797(graph, []int{}, 0)

	return ans797
}

func dfs797(graph [][]int, prefix []int, entry int) {
	if entry == len(graph)-1 {
		prefix = append(prefix, entry)
		ans797 = append(ans797, prefix)
		return
	}

	for i := 0; i < len(graph[entry]); i++ {
		pp := []int{}
		pp = append(pp, prefix...)
		pp = append(pp, entry)
		if graph[entry][i] == entry {
			continue
		}
		dfs797(graph, pp, graph[entry][i])
	}
}
