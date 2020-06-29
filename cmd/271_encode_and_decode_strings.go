package main

import "strings"

type Codec struct {
    length []int
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var builder strings.Builder
	codec.length = make([]int, len(strs))
	for k, v := range strs {
		builder.WriteString(v)
		codec.length[k] = len(v)
	}
	return builder.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	res := make([]string, len(codec.length))
    for k, v := range codec.length {
		res[k] , strs = strs[:v], strs[v:]
	}
	return res
}