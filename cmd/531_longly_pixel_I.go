package main

func findLonelyPixel(picture [][]byte) int {
	m := len(picture)
	n := len(picture[0])

	rows := make([]int, m)
	cols := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}

	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] == 1 && cols[j] == 1 && picture[i][j] == 'B' {
				count++
			}
		}
	}

	return count
}
