package main

import "fmt"

//lint:ignore U1000 unused
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	rowl, rowr := 0, m-1

	for rowl < rowr {
		mid := (rowl + rowr) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			rowr = mid - 1
		} else {
			if mid+1 < m && matrix[mid+1][0] > target {
				rowl = mid
				break
			}
			rowl = mid + 1
		}
	}

	row := rowl

	if matrix[row][0] == target {
		return true
	}

	coll, colr := 0, n-1

	for coll < colr {
		mid := (coll + colr) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			colr = mid - 1
		} else {
			coll = mid + 1
		}
	}

	fmt.Println(row, coll, colr)

	return matrix[row][coll] == target
}
