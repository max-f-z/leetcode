package main

import "sort"

//lint:ignore U1000 unused
func isReflected(points [][]int) bool {
	vals := map[int]map[int]bool{}
	nodups := [][]int{}
	for i, v := range points {
		if val, ok := vals[v[0]]; !ok {
			vals[v[0]] = make(map[int]bool)
		} else if val[v[1]] {
			continue
		}
		vals[v[0]][v[1]] = true
		nodups = append(nodups, points[i])
	}

	if len(nodups) == 1 {
		return true
	}

	sort.Slice(nodups, func(i, j int) bool {
		return nodups[i][0] < nodups[j][0]
	})

	var x float64
	if len(nodups)%2 == 0 {
		x = (float64(nodups[len(nodups)/2][0]) + float64(nodups[len(nodups)/2-1][0])) / 2
	} else {
		x = float64(nodups[len(nodups)/2][0])
	}

	for _, v := range nodups {
		if !vals[int(2*x)-v[0]][v[1]] {
			return false
		}
	}

	return true
}
