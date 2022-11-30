package main

func uniqueOccurrences(arr []int) bool {
	count := map[int]int{}

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	occurence := map[int]int{}

	for k, v := range count {
		if occurence[v] != 0 {
			return false
		}
		occurence[v] = k
	}

	return true
}
