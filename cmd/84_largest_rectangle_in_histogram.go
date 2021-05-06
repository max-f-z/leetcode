package main

import (
	"container/list"
)

func largestRectangleArea(heights []int) int {
	stack := list.New()

	maxArea := -1

	stack.PushBack(-1)

	for i := 0; i < len(heights); i++ {
		if stack.Len() != 1 {
			top := stack.Back().Value.(int)

			for stack.Len() > 1 && heights[i] < heights[top] {
				height := heights[top]

				stack.Remove(stack.Back())
				top = stack.Back().Value.(int)

				width := i - top - 1

				if height*width > maxArea {
					maxArea = height * width
				}
			}
		}

		stack.PushBack(i)
	}

	for stack.Len() > 1 {
		top := stack.Back().Value.(int)

		height := heights[top]

		stack.Remove(stack.Back())

		top = stack.Back().Value.(int)

		width := len(heights) - top - 1

		if height*width > maxArea {
			maxArea = height * width
		}

	}

	return maxArea
}
