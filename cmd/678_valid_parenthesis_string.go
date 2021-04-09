package main

import (
	"container/list"
)

func checkValidString(s string) bool {
	s1 := list.New()
	s2 := list.New()

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			s1.PushBack(i)
		} else if s[i] == '*' {
			s2.PushBack(i)
		} else {
			if s1.Len() > 0 {
				s1.Remove(s1.Back())
			} else if s2.Len() > 0 {
				s2.Remove(s2.Back())
			} else {
				return false
			}
		}
	}

	for s1.Len() > 0 {
		if s2.Len() > 0 {
			idx1 := s1.Remove(s1.Back()).(int)
			idx2 := s2.Remove(s2.Back()).(int)

			// idx1 ( idx2 *  顺序反了变成 )( 非法括号

			if idx1 > idx2 {
				return false
			}
			continue
		}
		return false
	}

	return true
}
