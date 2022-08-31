package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific := make([][]int, m)
	atlantic := make([][]int, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]int, n)
		atlantic[i] = make([]int, n)
	}

	pqueue := [][]int{}
	aqueue := [][]int{}

	for i := 0; i < m; i++ {
		pacific[i][0] = 1
		atlantic[i][n-1] = 1
		pqueue = append(pqueue, []int{i, 0})
		aqueue = append(aqueue, []int{i, n - 1})
	}

	for i := 0; i < n; i++ {
		pacific[0][i] = 1
		atlantic[m-1][i] = 1
		pqueue = append(pqueue, []int{0, i})
		aqueue = append(aqueue, []int{m - 1, i})
	}

	for len(pqueue) > 0 {
		newQ := [][]int{}
		for _, tuple := range pqueue {
			x, y := tuple[0], tuple[1]

			if x == 35 && y == 2 {
				fmt.Println(x, y)
			}

			if x+1 >= 0 && x+1 < m && heights[x+1][y] >= heights[x][y] && pacific[x+1][y] != 1 {
				pacific[x+1][y] = 1
				newQ = append(newQ, []int{x + 1, y})
			}

			if x-1 >= 0 && x-1 < m && heights[x-1][y] >= heights[x][y] && pacific[x-1][y] != 1 {
				pacific[x-1][y] = 1
				newQ = append(newQ, []int{x - 1, y})
			}

			if y+1 >= 0 && y+1 < n && heights[x][y+1] >= heights[x][y] && pacific[x][y+1] != 1 {
				pacific[x][y+1] = 1
				newQ = append(newQ, []int{x, y + 1})
			}

			if y-1 >= 0 && y-1 < n && heights[x][y-1] >= heights[x][y] && pacific[x][y-1] != 1 {
				pacific[x][y-1] = 1
				newQ = append(newQ, []int{x, y - 1})
			}
		}

		pqueue = newQ
	}

	for len(aqueue) > 0 {
		newQ := [][]int{}
		for _, tuple := range aqueue {
			x, y := tuple[0], tuple[1]

			if x+1 >= 0 && x+1 < m && heights[x+1][y] >= heights[x][y] && atlantic[x+1][y] != 1 {
				atlantic[x+1][y] = 1
				newQ = append(newQ, []int{x + 1, y})
			}

			if x-1 >= 0 && x-1 < m && heights[x-1][y] >= heights[x][y] && atlantic[x-1][y] != 1 {
				atlantic[x-1][y] = 1
				newQ = append(newQ, []int{x - 1, y})
			}

			if y+1 >= 0 && y+1 < n && heights[x][y+1] >= heights[x][y] && atlantic[x][y+1] != 1 {
				atlantic[x][y+1] = 1
				newQ = append(newQ, []int{x, y + 1})
			}

			if y-1 >= 0 && y-1 < n && heights[x][y-1] >= heights[x][y] && atlantic[x][y-1] != 1 {
				atlantic[x][y-1] = 1
				newQ = append(newQ, []int{x, y - 1})
			}
		}

		aqueue = newQ
	}

	results := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] == 1 && atlantic[i][j] == 1 {
				results = append(results, []int{i, j})
			}
		}
	}

	return results
}
