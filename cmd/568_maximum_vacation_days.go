package main

import (
	"math"
)

// if we start ith city jth week max vacation days
var memo [][]int

//lint:ignore U1000 unused
func maxVacationDays(flights [][]int, days [][]int) int {
	cities := len(days)
	weeks := len(days[0])

	memo = make([][]int, cities)
	for i := 0; i < cities; i++ {
		memo[i] = make([]int, weeks)
		for j := 0; j < weeks; j++ {
			memo[i][j] = math.MinInt32
		}
	}

	return maxVacationDaysMemoDFS(flights, days, 0, 0, cities, weeks)
}

func maxVacationDaysMemoDFS(flights [][]int, days [][]int, city int, week int, cities int, weeks int) int {
	max := 0

	if week == weeks {
		return 0
	}

	if memo[city][week] != math.MinInt32 {
		return memo[city][week]
	}

	for i := 0; i < cities; i++ {
		if flights[city][i] == 1 || i == city {
			res := days[i][week] + maxVacationDaysMemoDFS(flights, days, i, week+1, cities, weeks)
			if max < res {
				max = res
			}
		}
	}

	memo[city][week] = max
	return max
}

//lint:ignore U1000 unused
func maxVacationDaysTLE(flights [][]int, days [][]int) int {
	cities := len(days)
	weeks := len(days[0])

	return maxVacationDaysDFS(flights, days, 0, 0, 0, cities, weeks)
}

func maxVacationDaysDFS(flights [][]int, days [][]int, prev int, city int, week int, cities int, weeks int) int {
	max := 0

	if week == weeks {
		return prev
	}

	for i := 0; i < cities; i++ {
		if flights[city][i] == 1 || i == city {
			res := maxVacationDaysDFS(flights, days, days[i][week]+prev, i, week+1, cities, weeks)
			if max < res {
				max = res
			}
		}
	}

	if max < prev {
		return prev
	}
	return max
}
