package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	cnt := map[int]int{}

	for i := 0; i < len(hand); i++ {
		cnt[hand[i]]++
	}

	sort.Ints(hand)

	splitCnt := 0
	for i := 0; i < len(hand); i++ {
		if cnt[hand[i]] > 0 {
			cnt[hand[i]]--
			for j := 1; j < groupSize; j++ {
				if cnt[hand[i]+j] == 0 {
					return false
				}
				cnt[hand[i]+j]--
			}
			splitCnt++
		}
	}

	return splitCnt*groupSize == len(hand)
}
