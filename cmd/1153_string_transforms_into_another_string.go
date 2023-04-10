package main

// str1  a,b,c
// str2  b,c,a

// step1 a->d d,b,c step2 c->a d,b,a step3 b->c d,c,a step4 d->b b,c,a
//
//	b,c,a            b,c,a            b,c,a            b,c,a
//
//lint:ignore U1000 unused
func canConvert(str1 string, str2 string) bool {
	cols := map[byte]byte{}

	if str1 == str2 {
		return true
	}

	l := len(str1)

	for i := 0; i < l; i++ {
		if _, ok := cols[str1[i]]; ok {
			if cols[str1[i]] == str2[i] {
				continue
			}
			return false
		}

		if str1[i] == str2[i] {
			cols[str1[i]] = str2[i]
			continue
		}

		cols[str1[i]] = str2[i]
	}

	buf := map[byte]bool{}

	for i := 0; i < l; i++ {
		buf[str2[i]] = true
	}

	return len(buf) != 26
}
