package main

import "testing"

func Test990(t *testing.T) {
	equaltions := []string{"a==b", "b!=a"}

	ans := equationsPossible(equaltions)

	t.Log(ans)
}
