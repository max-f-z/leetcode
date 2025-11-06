package main

import "container/heap"

func processQueries(c int, connections [][]int, queries [][]int) []int {
	// Build union-find structure to identify connected components
	unf := newUF(c)
	for _, conn := range connections {
		unf.Union(conn[0]-1, conn[1]-1)
	}

	// Map each station to its component root
	roots := make([]int, c)
	for i := 0; i < c; i++ {
		roots[i] = unf.Find(i)
	}

	// Create priority queue for each component with all its stations
	componentPQs := buildComponentPriorityQueues(c, roots)

	// Process queries
	return executeQueries(queries, roots, componentPQs)
}

// buildComponentPriorityQueues creates a min-heap for each component containing all station IDs
func buildComponentPriorityQueues(c int, roots []int) map[int]*pq3607 {
	pqs := map[int]*pq3607{}

	for i := 0; i < c; i++ {
		root := roots[i]
		stationID := i + 1

		if pq, exists := pqs[root]; exists {
			heap.Push(pq, stationID)
		} else {
			pq := &pq3607{}
			heap.Init(pq)
			heap.Push(pq, stationID)
			pqs[root] = pq
		}
	}

	return pqs
}

// executeQueries processes all queries and returns results for type-1 queries
func executeQueries(queries [][]int, roots []int, componentPQs map[int]*pq3607) []int {
	const (
		queryTypeGet     = 1
		queryTypeOffline = 2
	)

	res := []int{}
	offlines := map[int]bool{}

	for _, query := range queries {
		queryType := query[0]
		station := query[1]

		if queryType == queryTypeGet {
			power := findAvailablePower(station, roots, componentPQs, offlines)
			res = append(res, power)
		} else {
			offlines[station] = true
		}
	}

	return res
}

// findAvailablePower returns the smallest available station in the component
// If the queried station is online, return it; otherwise find the next smallest available
func findAvailablePower(station int, roots []int, componentPQs map[int]*pq3607, offlines map[int]bool) int {
	// If station is online, it can power itself
	if !offlines[station] {
		return station
	}

	// Station is offline, find smallest available in component
	componentRoot := roots[station-1]
	pq := componentPQs[componentRoot]

	return getSmallestAvailableStation(pq, offlines)
}

// getSmallestAvailableStation pops offline stations until finding an online one
func getSmallestAvailableStation(pq *pq3607, offlines map[int]bool) int {
	if pq.Len() == 0 {
		return -1
	}

	// Keep popping offline stations until we find an online one
	for pq.Len() > 0 {
		top := pq.Top()

		if !offlines[top] {
			return top
		}

		heap.Pop(pq)
	}

	return -1
}

type uf struct {
	parent     []int
	ranks      []int
	components int
}

func newUF(n int) *uf {
	p, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &uf{
		parent:     p,
		ranks:      r,
		components: n,
	}
}

func (u *uf) Find(x int) int {
	if x == u.parent[x] {
		return x
	}

	p := u.Find(u.parent[x])
	u.parent[x] = p
	return p
}

func (u *uf) Union(x, y int) {
	xx, yy := u.Find(x), u.Find(y)

	if xx == yy {
		return
	}

	if u.ranks[xx] < u.ranks[yy] {
		u.parent[xx] = yy
	} else if u.ranks[xx] > u.ranks[yy] {
		u.parent[yy] = xx
	} else {
		u.parent[xx] = yy
		u.ranks[yy]++
	}

	u.components--
}

type pq3607 []int

func (pq pq3607) Len() int {
	return len(pq)
}

func (pq pq3607) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq pq3607) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pq3607) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *pq3607) Pop() (v interface{}) {
	v, *pq = (*pq)[pq.Len()-1], (*pq)[:pq.Len()-1]
	return v
}

func (pq *pq3607) Top() int {
	return (*pq)[0]
}
