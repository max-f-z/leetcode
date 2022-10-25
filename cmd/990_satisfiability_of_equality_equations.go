package main

type UF struct {
	count  int
	parent []int
}

func createLowerUF() *UF {
	arr := make([]int, 26)

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	return &UF{
		count:  26,
		parent: arr,
	}
}

func (uf *UF) Find(x int) int {
	if uf.parent[x] == x {
		return x
	}

	p := uf.Find(uf.parent[x])
	uf.parent[x] = p
	return p
}

func (uf *UF) Union(x, y int) {
	xx, yy := uf.Find(x), uf.Find(y)

	if xx != yy {
		uf.parent[yy] = xx
		uf.count--
	}
}

func equationsPossible(equations []string) bool {
	uf := createLowerUF()

	for _, eq := range equations {
		if eq[1] == '=' {
			x := int(eq[0] - 'a')
			y := int(eq[3] - 'a')
			uf.Union(x, y)
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' {
			x := int(eq[0] - 'a')
			y := int(eq[3] - 'a')
			xx, yy := uf.Find(x), uf.Find(y)

			if xx == yy {
				return false
			}
		}
	}

	return true
}
