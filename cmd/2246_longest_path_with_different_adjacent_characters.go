package main

func longestPath(parent []int, s string) int {
	child := map[int][]int{}

	for i, v := range parent {
		if i == 0 {
			continue
		}
		if child[v] == nil {
			child[v] = []int{}
		}
		child[v] = append(child[v], i)
	}

	p := &path{
		max:   0,
		child: child,
		s:     s,
	}

	p.longestPathDFS(0)

	return p.max
}

type path struct {
	max   int
	child map[int][]int
	s     string
}

func (p *path) longestPathDFS(cur int) int {
	longest := 0
	secondLongest := 0

	if p.child[cur] == nil || len(p.child[cur]) == 0 {
		if p.max < 1 {
			p.max = 1
		}
		return 1
	}

	for _, c := range p.child[cur] {
		val := p.longestPathDFS(c)
		if p.s[c] == p.s[cur] {
			continue
		}
		if val > longest {
			secondLongest = longest
			longest = val
		} else if val > secondLongest {
			secondLongest = val
		}
	}

	if secondLongest+longest+1 > p.max {
		p.max = secondLongest + longest + 1
	}

	if longest != 0 {
		return 1 + longest
	}

	if p.max < 1 {
		p.max = 1
	}

	return 1
}
