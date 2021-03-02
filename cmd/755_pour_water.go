package main

func pourWater(heights []int, V int, K int) []int {
	for i := 0; i < V; i++ {
		h := heights[K] + 1

		// fill left
		left := false
		idx := -1
		if K > 0 {
			for j := K - 1; j >= 0; j-- {
				if heights[j]+1 < h {
					h = heights[j] + 1
					idx = j
					left = true
				}

				if heights[j]+1 > h {
					break
				}
			}

			if left {
				heights[idx] = heights[idx] + 1
				continue
			}
		}

		// fill right
		right := false
		if K < len(heights)-1 {
			for j := K + 1; j < len(heights); j++ {
				if heights[j]+1 < h {
					h = heights[j] + 1
					idx = j
					right = true
				}

				if heights[j]+1 > h {
					break
				}
			}
			if right {
				heights[idx] = heights[idx] + 1
				continue
			}
		}

		heights[K] = heights[K] + 1
	}

	return heights
}
