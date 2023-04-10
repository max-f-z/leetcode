package main

type UnionFind323 struct {
	count  int
	parent []int
}

func (u *UnionFind323) Find(x int) int {
	if u.parent[x] == x {
		return x
	}

	p := u.Find(u.parent[x])
	u.parent[x] = p
	return p
}

func (u *UnionFind323) Union(x, y int) {
	xx, yy := u.Find(x), u.Find(y)

	if xx != yy {
		u.parent[xx] = yy
		u.count--
	}
}

func (u *UnionFind323) Count() int {
	return u.count
}

func CreateUnionFind323(n int) *UnionFind323 {
	uf := &UnionFind323{}
	uf.count = n
	uf.parent = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

//lint:ignore U1000 unused
func countComponents(n int, edges [][]int) int {
	uf := CreateUnionFind323(n)

	for _, v := range edges {
		uf.Union(v[0], v[1])
	}

	return uf.count
}
