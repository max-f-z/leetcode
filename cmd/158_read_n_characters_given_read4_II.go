package main

//lint:ignore U1000 unused
var solution1 = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	pos := 0
	size := 0
	buffer := make([]byte, 4)
	return func(buf []byte, n int) int {
		cnt := 0
		for cnt < n {
			if size == pos {
				pos = 0
				size = read4(buffer)
				if size == 0 {
					return cnt
				}
			}
			buf[cnt] = buffer[pos]
			cnt++
			pos++
		}
		return cnt
	}
}
