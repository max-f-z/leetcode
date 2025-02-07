package main

func queryResults(limit int, queries [][]int) []int {
	refBallColor := make(map[int]int)
	refColorBall := make(map[int]int)

	results := make([]int, len(queries))

	for idx, query := range queries {
		position, color := query[0], query[1]

		if refBallColor[position] != 0 {
			oldColor := refBallColor[position]
			refColorBall[oldColor]--

			if refColorBall[oldColor] == 0 {
				delete(refColorBall, oldColor)
			}
		}

		refBallColor[position] = color
		refColorBall[color]++

		results[idx] = len(refColorBall)
	}

	return results
}
