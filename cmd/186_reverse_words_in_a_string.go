package main

//lint:ignore U1000 unused
func reverseWords(s []byte) {
	l := 0
	h := len(s) - 1
	for l < h {
		s[l], s[h] = s[h], s[l]
		l++
		h--
	}

	start := 0
	end := 0
	for i, v := range s {
		if v != ' ' {
			end++
		} else {
			reverseByte(s, start, end-1)
			start = i + 1
			end = start
		}
	}
	reverseByte(s, start, end-1)
}

func reverseByte(s []byte, l, h int) {
	for l < h {
		s[l], s[h] = s[h], s[l]
		l++
		h--
	}
}
