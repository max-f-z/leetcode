package main

func numberOfWays(corridor string) int {
	cnt := 0

	for i := 0; i < len(corridor); i++ {
		if corridor[i] == 'S' {
			cnt++
		}
	}

	if cnt%2 == 1 || cnt <= 0 {
		return 0
	}

	if cnt == 2 {
		return 1
	}

	cnt = 0
	ans := 1
	prev := 0
	mod := 1000000007
	for i := 0; i < len(corridor); i++ {
		if corridor[i] == 'S' {
			cnt++
			if cnt == 3 {
				ans = (ans * (i - prev)) % mod
				cnt = 1
			}
			prev = i
		}
	}

	return ans
}
