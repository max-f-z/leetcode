package main

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {

		for i := 0; i < n; i++ {
			skip := false
			for j := 0; j < n; j++ {
				if j != i {
					if !knows(j, i) || knows(i, j) {
						skip = true
						break
					}
				}
			}
			if !skip {
				return i
			}
		}
		return -1
	}
}
