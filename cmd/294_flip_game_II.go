package main

//lint:ignore U1000 unused
func canWin(s string) bool {
	return canWinPlayer1(s)
}

func canWinPlayer1(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		tmp := []byte(s)
		if s[i] == '+' && s[i+1] == '+' {
			tmp[i], tmp[i+1] = 45, 45
			s = string(tmp)

			if !canWinPlayer2(s) {
				return true
			}
			tmp[i], tmp[i+1] = 43, 43
			s = string(tmp)
		}
	}
	return false
}

func canWinPlayer2(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		tmp := []byte(s)
		if s[i] == '+' && s[i+1] == '+' {
			tmp[i], tmp[i+1] = 45, 45
			s = string(tmp)

			if !canWinPlayer1(s) {
				return true
			}
			tmp[i], tmp[i+1] = 43, 43
			s = string(tmp)
		}
	}
	return false
}
