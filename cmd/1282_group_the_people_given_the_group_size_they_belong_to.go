package main

func groupThePeople(groupSizes []int) [][]int {
	n := len(groupSizes)

	ref := map[int][]int{}

	for i := 0; i < n; i++ {
		if _, ok := ref[groupSizes[i]]; !ok {
			ref[groupSizes[i]] = []int{}
		}

		ref[groupSizes[i]] = append(ref[groupSizes[i]], i)
	}

	ans := [][]int{}

	for k, v := range ref {
		for i := 0; i < len(v); i += k {
			tmp := make([]int, k)
			for j := i; j-i < k; j++ {
				tmp[j-i] = v[j]
			}
			ans = append(ans, tmp)
		}
	}

	return ans
}
