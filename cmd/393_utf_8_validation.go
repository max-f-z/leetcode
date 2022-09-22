package main

import "strconv"

func validUtf8(data []int) bool {
	for len(data) > 0 {
		bits := strconv.FormatInt(int64(data[0]), 2)
		data = data[1:]
		leadingOnes := 0

		for _, bit := range bits {
			if bit == '0' {
				break
			}

			leadingOnes++
		}
		if len(bits) != 8 {
			continue
		}

		if leadingOnes == 1 || leadingOnes > 4 {
			return false
		}

		for i := 1; i < leadingOnes; i++ {
			if len(data) == 0 {
				return false
			}

			bits = strconv.FormatInt(int64(data[0]), 2)
			data = data[1:]

			if len(bits) < 8 || bits[1] != '0' {
				return false
			}
		}
	}

	return true
}
