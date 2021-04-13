package main

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
import (
	"container/list"
)

type cache struct {
	l        *list.List
	cache    map[int]*list.Element
	capacity int
	length   int
}

func (c *cache) set(x, y int) {
	if el, ok := c.cache[x]; ok {
		c.l.MoveToFront(el)
	} else {
		c.l.PushFront(x)
		c.cache[x] = c.l.Front()
		c.length++

		if c.length > c.capacity {
			c.length--
			key := c.l.Back().Value.(int)
			c.l.Remove(c.cache[key])
			delete(c.cache, key)
		}
	}
}

func (c *cache) get(x int) int {
	if _, ok := c.cache[x]; !ok {
		return -1
	}
	c.l.MoveToFront(c.cache[x])
	return c.cache[x].Value.(int)
}

func LRU(operators [][]int, k int) []int {
	// write code here
	c := &cache{
		l:        list.New(),
		cache:    map[int]*list.Element{},
		capacity: k,
	}

	res := []int{}
	for _, v := range operators {
		if v[0] == 1 {
			c.set(v[1], v[2])
		} else if v[0] == 2 {
			vv := c.get(v[1])
			res = append(res, vv)
		}
	}
	return res
}
