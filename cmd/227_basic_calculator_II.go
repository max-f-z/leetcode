package main

import (
	"container/list"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
func calculate227(s string) int {
	s = strings.ReplaceAll(s, " ", "")

	nums := list.New().Init()
	ops := list.New().Init()

	prefix := strings.Builder{}
	tmp := -1
	var tmpOp byte
	tmpOp = '%'
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			val, _ := strconv.Atoi(prefix.String())
			if tmp == -1 {
				nums.PushBack(val)
			} else {
				if tmpOp == '*' {
					nums.PushBack(val * tmp)
				} else {
					nums.PushBack(tmp / val)
				}
				tmp = -1
				tmpOp = '%'
			}
			ops.PushBack(s[i])
			prefix.Reset()
		} else if s[i] == '*' || s[i] == '/' {
			val, _ := strconv.Atoi(prefix.String())
			if tmp == -1 {
				tmp = val
			} else {
				if tmpOp == '*' {
					tmp = val * tmp
				} else {
					tmp = tmp / val
				}
			}
			tmpOp = s[i]
			prefix.Reset()
		} else {
			prefix.WriteByte(s[i])
		}
	}

	val, _ := strconv.Atoi(prefix.String())
	if tmp == -1 {
		nums.PushBack(val)
	} else {
		if tmpOp == '*' {
			nums.PushBack(val * tmp)
		} else {
			nums.PushBack(tmp / val)
		}
	}

	res := nums.Front().Value.(int)
	nums.Remove(nums.Front())

	for ops.Len() > 0 {
		op := ops.Front().Value.(byte)
		val := nums.Front().Value.(int)
		ops.Remove(ops.Front())
		nums.Remove(nums.Front())
		if op == '+' {
			res += val
		} else {
			res -= val
		}
	}

	return res
}
