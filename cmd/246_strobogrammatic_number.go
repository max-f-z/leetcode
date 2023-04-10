package main

//lint:ignore U1000 unused
func isStrobogrammatic(num string) bool {
	dict := map[string]byte{
		"0": '0',
		"1": '1',
		"6": '9',
		"8": '8',
		"9": '6',
	}
	convert := []byte(num)
	for i, v := range convert {
		if dict[string(v)] == 0 {
			return false
		}
		convert[i] = dict[string(v)]
	}

	l, h := 0, len(convert)-1
	for l < h {
		convert[l], convert[h] = convert[h], convert[l]
		l++
		h--
	}

	return string(convert) == num
}
