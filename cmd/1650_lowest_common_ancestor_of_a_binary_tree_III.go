package main

type Node1650 struct {
	Val    int
	Left   *Node1650
	Right  *Node1650
	Parent *Node1650
}

func lowestCommonAncestor1650(p *Node1650, q *Node1650) *Node1650 {
	ancestorsOfP := map[*Node1650]bool{}

	if p == q {
		return p
	}

	for p != nil {
		ancestorsOfP[p] = true
		p = p.Parent
	}

	for q != nil {
		if ancestorsOfP[q] {
			return q
		}
		q = q.Parent
	}

	return nil
}
