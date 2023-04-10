package main

//lint:ignore U1000 unused
func findBlackPixel(picture [][]byte, N int) int {
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

	record := map[int](map[int]bool){}

	for i := 0; i < m; i++ {
		if rows[i] != N {
			continue
		}
		for j := 0; j < n; j++ {
			if cols[j] != N {
				continue
			}

			if record[i][j] {
				continue
			}

			arr := []int{}
			for u := 0; u < m; u++ {
				if picture[u][j] == 'B' {
					arr = append(arr, u)
				}
			}

			valid := true
			for u := 1; u < N; u++ {
				if !identicalByteArray(picture[arr[0]], picture[arr[u]]) {
					valid = false
					break
				}
			}

			if valid {
				for u := 0; u < N; u++ {
					if record[arr[u]] == nil {
						record[arr[u]] = map[int]bool{}
					}
					record[arr[u]][j] = true
				}
			}
		}
	}

	count := 0
	for k := range record {
		count += len(record[k])
	}

	return count
}

func identicalByteArray(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
