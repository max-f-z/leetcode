package main

import "testing"

func Test588(t *testing.T) {
	fs := Constructor588()

	fs.Ls("/")

	fs.Mkdir("/a/b/c")

	fs.Ls("/a/b")
}
