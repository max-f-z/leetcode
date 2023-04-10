package main

import "sort"

type profitInfo struct {
	start  int
	end    int
	profit int
}

type profits []*profitInfo

func (p profits) Less(i, j int) bool {
	return p[i].end < p[j].end
}

func (p profits) Len() int {
	return len(p)
}

func (p profits) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//lint:ignore U1000 unused
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	ps := profits{}
	for i := 0; i < len(startTime); i++ {
		p := &profitInfo{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
		ps = append(ps, p)
	}

	sort.Sort(ps)

	prev := make([]int, ps.Len()+1)
	dp := make([]int, ps.Len()+1)
	prev[0] = 0
	dp[0] = 0

	for i := 1; i < ps.Len()+1; i++ {
		for j := i - 1; j >= 0; j-- {
			if ps[j].end <= ps[i-1].start {
				prev[i] = j + 1
				break
			}
		}
	}

	dp[1] = ps[0].profit

	for i := 2; i <= ps.Len(); i++ {
		if dp[i-1] < dp[prev[i]]+ps[i-1].profit {
			dp[i] = dp[prev[i]] + ps[i-1].profit
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[ps.Len()]
}
