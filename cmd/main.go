package main

import (
	"fmt"
)

func main() {
	obstacleGrid := [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
