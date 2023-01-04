package main

import "testing"

func Test2244(t *testing.T) {
	tasks := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}

	ans := minimumRounds(tasks)

	t.Log(ans)
}
