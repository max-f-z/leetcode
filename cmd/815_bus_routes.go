package main

func numBusesToDestination(routes [][]int, source int, target int) int {
	// key - bus no.  value - station
	busStation := map[int]map[int]bool{}

	// key - station  value - bus no.
	stationBus := map[int]map[int]bool{}

	for k, v := range routes {
		if busStation[k] == nil {
			busStation[k] = make(map[int]bool)
		}

		for _, s := range v {
			if stationBus[s] == nil {
				stationBus[s] = make(map[int]bool)
			}

			busStation[k][s] = true
			stationBus[s][k] = true
		}
	}

	seenBus := map[int]bool{}
	seenStation := map[int]bool{}

	queue := []int{source}
	res := 1
	seenStation[source] = true

	for len(queue) > 0 {
		newQ := []int{}
		for _, v := range queue {
			if v == target {
				return res - 1
			}

			bs := stationBus[v]
			for bk := range bs {
				if !seenBus[bk] {
					seenBus[bk] = true
				} else {
					continue
				}

				for sk := range busStation[bk] {
					if !seenStation[sk] {
						seenStation[sk] = true
						newQ = append(newQ, sk)
					}
				}
			}

		}
		res++
		queue = newQ
	}

	return -1
}
