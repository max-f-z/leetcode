package main

import "testing"

func Test106(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	ans := buildTree106(inorder, postorder)

	t.Log(ans)
}
