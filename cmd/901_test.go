package main

import "testing"

func Test901(t *testing.T) {
	ss := Constructor901()

	ans := []int{}

	ans = append(ans, ss.Next(100))
	ans = append(ans, ss.Next(80))
	ans = append(ans, ss.Next(60))
	ans = append(ans, ss.Next(70))
	ans = append(ans, ss.Next(60))
	ans = append(ans, ss.Next(75))
	ans = append(ans, ss.Next(85))

	t.Log(ans)
}
