package main

import (
	"container/heap"
	"sort"
)

type Point struct {
	row int
	col int
	val int
}

type pqPoint []*Point

func (p pqPoint) Len() int {
	return len(p)
}

func (p pqPoint) Less(i, j int) bool {
	return p[i].val < p[j].val
}

func (p pqPoint) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pqPoint) Push(x interface{}) {
	*p = append(*p, x.(*Point))
}

func (p *pqPoint) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p *pqPoint) Top() Point {
	return *(*p)[0]
}

func maxPoints(grid [][]int, queries []int) []int {
	sortedQueries := make([]int, len(queries))
	copy(sortedQueries, queries)
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i] < sortedQueries[j]
	})

	ref := map[int]int{}

	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	pq := &pqPoint{}
	heap.Init(pq)
	heap.Push(pq, &Point{row: 0, col: 0, val: grid[0][0]})
	visited[0][0] = true

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	points := 0

	for i := 0; i < len(sortedQueries); i++ {
		query := sortedQueries[i]

		for pq.Len() > 0 && pq.Top().val < query {
			current := heap.Pop(pq).(*Point)
			points += 1

			for _, d := range directions {
				newRow := current.row + d[0]
				newCol := current.col + d[1]

				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] {
					heap.Push(pq, &Point{row: newRow, col: newCol, val: grid[newRow][newCol]})
					visited[newRow][newCol] = true
				}
			}
		}

		ref[query] = points
	}

	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = ref[queries[i]]
	}

	return res
}
