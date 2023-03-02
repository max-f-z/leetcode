package main

import "testing"

func Test443(t *testing.T) {
	cases := []byte{'a', 'a', 'a', 'a', 'b', 'c'}

	ans := compress(cases)

	t.Log(ans)
}
