package main

func removeStones(stones [][]int) int {
	parents := make([]int, len(stones))
	for i := 0; i < len(stones); i++ {
		parents[i] = i
	}
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				pi, pj := findParent(parents, i), findParent(parents, j)

				parents[pj] = pi
			}
		}
	}
	res := 0
	for i := 0; i < len(stones); i++ {
		if parents[i] == i {
			res++
		}
	}

	return len(stones) - res
}

func findParent(parent []int, idx int) int {
	i := idx
	for parent[i] != i {
		i = parent[i]
	}
	parent[idx] = i
	return i
}
