package main

import "sort"

func frequencySort(nums []int) []int {
	freq := map[int]int{}

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	type tuple struct {
		Num  int
		Freq int
	}

	tuples := make([]*tuple, len(nums))
	for i := 0; i < len(nums); i++ {
		tuples[i] = &tuple{
			Num:  nums[i],
			Freq: freq[nums[i]],
		}
	}

	sort.Slice(tuples, func(i, j int) bool {
		if tuples[i].Freq == tuples[j].Freq {
			return tuples[i].Num > tuples[j].Num
		}
		return tuples[i].Freq < tuples[j].Freq
	})

	for i := 0; i < len(nums); i++ {
		nums[i] = tuples[i].Num
	}

	return nums
}
