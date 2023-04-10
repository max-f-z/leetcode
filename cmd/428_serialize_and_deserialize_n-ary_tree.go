package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node428 struct {
	Val      int
	Children []*Node428
}

type Codec428 struct {
}

func Constructor428() *Codec428 {
	codec := Codec428{}
	return &codec
}

//lint:ignore ST1006 this
//lint:ignore U1000 unused
func (this *Codec428) serialize(root *Node428) string {
	return serializeHelper(root)
}

//lint:ignore ST1006 this
//lint:ignore U1000 unused
func (this *Codec428) deserialize(data string) *Node428 {
	return deserializeHelper(data)
}

//lint:ignore U1000 unused
func serializeHelper(node *Node428) string {
	if node == nil {
		return ""
	}

	if node.Children == nil || len(node.Children) == 0 {
		return strconv.Itoa(node.Val)
	}

	strs := []string{}
	for i := 0; i < len(node.Children); i++ {
		strs = append(strs, serializeHelper(node.Children[i]))
	}

	return fmt.Sprintf("%d[%s]", node.Val, strings.Join(strs, " "))
}

//lint:ignore U1000 unused
func deserializeHelper(data string) *Node428 {
	if data == "" {
		return nil
	}

	if !strings.Contains(data, "[") {
		val, _ := strconv.Atoi(data)
		return &Node428{Val: val}
	}

	idx := strings.Index(data, "[")

	val, _ := strconv.Atoi(data[:idx])

	tmp := &Node428{Val: val}
	tmp.Children = []*Node428{}

	sub := data[idx+1 : len(data)-1]
	flag := 0
	start := 0
	for i := 0; i < len(sub); i++ {
		if sub[i] == '[' {
			flag++
			continue
		}

		if sub[i] == ']' {
			flag--
			continue
		}

		if sub[i] == ' ' && flag == 0 {
			node := deserializeHelper(sub[start:i])
			tmp.Children = append(tmp.Children, node)
			start = i + 1
			continue
		}

	}

	if start <= len(sub) {
		str := sub[start:]
		node := deserializeHelper(str)
		tmp.Children = append(tmp.Children, node)
	}

	return tmp
}
