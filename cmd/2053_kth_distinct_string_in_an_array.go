package main

func kthDistinct(arr []string, k int) string {
	cnt := map[string]int{}

	for i := 0; i < len(arr); i++ {
		cnt[arr[i]]++
	}

	ans := ""
	idx := 0
	for i := 0; i < len(arr); i++ {
		if cnt[arr[i]] == 1 {
			idx++
		}
		if idx == k {
			ans = arr[i]
			break
		}
	}

	return ans
}
