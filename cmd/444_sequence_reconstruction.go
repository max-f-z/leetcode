package main

// type topoGraph struct {
// 	parent map[int][]int
// }

func sequenceReconstruction(org []int, seqs [][]int) bool {
	g := make(map[int][]int)
	degree := make(map[int]int)
	for _, v := range seqs {
		if len(v) == 0 {
			continue
		}
		start := v[len(v)-1]
		if degree[start] == 0 {
			degree[start] = 0
		}
		for j := len(v) - 2; j >= 0; j-- {
			g[start] = append(g[start], v[j])
			degree[v[j]]++
			start = v[j]

		}
	}
	var next []int
	for i, v := range degree {
		if v == 0 {
			next = append(next, i)
		}
	}
	var out []int
	for len(next) == 1 {
		c := next[0]
		next = next[1:]
		out = append(out, c)
		for _, n := range g[c] {
			degree[n]--
			if degree[n] == 0 {
				next = append(next, n)
			}
		}
		delete(g, c)
	}
	if len(next) != 0 || len(out) != len(org) || len(g) > 0 {
		return false
	}
	for i := 0; i < len(org); i++ {
		if org[i] != out[len(out)-1-i] {
			return false
		}
	}
	return true
}
