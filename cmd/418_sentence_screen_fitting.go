package main

func wordsTyping(sentence []string, rows int, cols int) int {
	row := 0
	col := 0

	for i := 0; i < rows; {
		for _, v := range sentence {
			tmp := len(v)
			if tmp > cols {
				return 0
			}
			if col+tmp <= cols {
				col = col + tmp + 1
				continue
			}
			col = tmp + 1
			i++
		}
		if i < rows {
			row++
		}
	}

	return row
}
