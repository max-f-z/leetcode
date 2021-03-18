package main

import "container/heap"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	mh := &maxHeap{}
	heap.Init(mh)
	ans := 0
	tank := startFuel
	prev := 0
	for i := 0; i < len(stations); i++ {
		tank -= (stations[i][0] - prev)
		prev = stations[i][0]

		for mh.Len() > 0 && tank < 0 {
			fuel := heap.Pop(mh).(int)
			tank += fuel
			ans++
		}

		if tank < 0 {
			return -1
		}

		heap.Push(mh, stations[i][1])
	}

	// see if the target can be reached
	{
		tank -= target - prev
		for mh.Len() > 0 && tank < 0 {
			tank += heap.Pop(mh).(int)
			ans++
		}

		if tank < 0 {
			return -1
		}
	}

	return ans
}
