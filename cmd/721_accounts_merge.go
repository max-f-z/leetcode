package main

import "sort"

//lint:ignore U1000 unused
func accountsMerge(accounts [][]string) [][]string {
	parents := make([]int, len(accounts))

	emails := make(map[string]int)

	for i, group := range accounts {
		parents[i] = i

		for _, email := range group[1:] {
			if _, ok := emails[email]; ok {
				parents[accountsMergeFindParent(parents, i)] = accountsMergeFindParent(parents, emails[email])
				continue
			}
			emails[email] = i
		}
	}

	merged := make(map[int][]string)

	for email, acctIdx := range emails {
		group := accountsMergeFindParent(parents, acctIdx)

		merged[group] = append(merged[group], email)
	}

	ans := make([][]string, 0)

	for group, emails := range merged {
		account := accounts[group][0]
		sort.Strings(emails)
		record := append([]string{account}, emails...)
		ans = append(ans, record)
	}

	return ans
}

//lint:ignore U1000 unused
func accountsMergeFindParent(parents []int, idx int) int {
	if parents[idx] == idx {
		return parents[idx]
	}

	parents[idx] = accountsMergeFindParent(parents, parents[idx])
	return parents[idx]
}
