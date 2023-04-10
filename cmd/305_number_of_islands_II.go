package main

type UnionFind struct {
	parent []int
	count  int
}

func CreateUnionFind(m, n int) *UnionFind {
	p := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		p[i] = i
	}

	return &UnionFind{
		parent: p,
		count:  0,
	}
}

func (u *UnionFind) Find(x int) int {
	if x == u.parent[x] {
		return x
	}

	p := u.Find(u.parent[x])
	u.parent[x] = p
	return p
}

func (u *UnionFind) AddIsland() {
	u.count++
}

func (u *UnionFind) Union(x, y int) {
	xx, yy := u.Find(x), u.Find(y)

	if xx != yy {
		u.parent[xx] = yy
		u.count--
	}

}

func (u *UnionFind) Count() int {
	return u.count
}

//lint:ignore U1000 unused
func numIslands2(m int, n int, positions [][]int) []int {
	uf := CreateUnionFind(m, n)
	islands := make([][]bool, m)
	for i := 0; i < m; i++ {
		islands[i] = make([]bool, n)
	}
	res := make([]int, len(positions))

	for i := 0; i < len(positions); i++ {
		x, y := positions[i][0], positions[i][1]

		if islands[x][y] {
			res[i] = uf.Count()
			continue
		}
		islands[x][y] = true

		uf.AddIsland()
		for j := 0; j < len(directions); j++ {
			xx, yy := x+directions[j][0], y+directions[j][1]
			if xx < 0 || yy < 0 || xx >= m || yy >= n || !islands[xx][yy] {
				continue
			}

			uf.Union(x*n+y, xx*n+yy)
		}
		res[i] = uf.Count()
	}

	return res
}

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}
