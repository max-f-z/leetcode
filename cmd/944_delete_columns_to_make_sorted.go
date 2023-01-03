package main

func minDeletionSize(strs []string) int {
	l := len(strs[0])
	n := len(strs)

	cnt := 0
	for i := 0; i < l; i++ {
		for j := 0; j < n-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				cnt++
				break
			}
		}
	}

	return cnt
}
