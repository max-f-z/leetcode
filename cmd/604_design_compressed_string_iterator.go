package main

type StringIterator struct {
	chars  []byte
	nums   []int
	prefix []int
	idx    int
}

func Constructor604(compressedString string) StringIterator {
	chars := []byte{}
	nums := []int{}
	prefix := []int{}

	var c byte
	cnt := 0
	tmp := 0

	c = compressedString[0]

	for i := 1; i < len(compressedString); i++ {
		if compressedString[i] >= '0' && compressedString[i] <= '9' {
			cnt = cnt*10 + int(compressedString[i]-48)
		} else {
			prefix = append(prefix, tmp)
			chars = append(chars, c)
			nums = append(nums, cnt)
			tmp += cnt
			c = compressedString[i]
			cnt = 0
		}
	}

	if cnt != 0 {
		prefix = append(prefix, tmp)
		chars = append(chars, c)
		nums = append(nums, cnt)
	}

	s := StringIterator{
		chars:  chars,
		nums:   nums,
		prefix: prefix,
		idx:    0,
	}

	return s
}

//lint:ignore ST1006 this
func (this *StringIterator) Next() byte {
	if this.idx < this.prefix[len(this.prefix)-1]+this.nums[len(this.nums)-1] {
		for i := range this.prefix {
			if this.idx < this.prefix[i]+this.nums[i] {
				this.idx++
				return this.chars[i]
			}
		}
	}
	return ' '
}

//lint:ignore ST1006 this
func (this *StringIterator) HasNext() bool {
	return this.idx < this.prefix[len(this.prefix)-1]+this.nums[len(this.nums)-1]
}
