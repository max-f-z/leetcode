package main

func generatePossibleNextMoves(s string) []string {
	res := []string{}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '+' && s[i+1] == '+' {
			tmp := []byte(s)
			tmp[i], tmp[i+1] = 45, 45
			res = append(res, string(tmp))
		}
	}

	return res
}
