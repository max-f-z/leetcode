package main

import "container/list"

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0

	m := len(matrix)
	n := 0
	if len(matrix) > 0 {
		n = len(matrix[0])
	}

	rows := [][]int{}

	for i := 0; i < m; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i != 0 {
					row[j] = rows[i-1][j] + 1
				} else {
					row[j] = 1
				}
			} else {
				row[j] = 0
			}
		}

		rows = append(rows, row)
	}

	for i := 0; i < m; i++ {
		local := largestRectangle(rows[i])

		if local > maxArea {
			maxArea = local
		}
	}

	return maxArea
}

func largestRectangle(heights []int) int {
	stack := list.New()

	stack.PushBack(-1)

	max := 0

	for i := 0; i < len(heights); i++ {
		if stack.Len() != 1 {
			top := stack.Back().Value.(int)

			for stack.Len() > 1 && heights[i] < heights[top] {
				h := heights[top]

				stack.Remove(stack.Back())
				top = stack.Back().Value.(int)

				w := i - 1 - top

				if h*w > max {
					max = h * w
				}
			}

		}

		stack.PushBack(i)
	}

	for stack.Len() > 1 {
		top := stack.Back().Value.(int)

		h := heights[top]

		stack.Remove(stack.Back())

		top = stack.Back().Value.(int)

		w := len(heights) - top - 1

		if h*w > max {
			max = h * w
		}
	}

	return max
}
