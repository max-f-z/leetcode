package main

//lint:ignore U1000 unused
func wordPatternMatch(pattern string, str string) bool {
	keys := map[string]string{}
	vals := map[string]string{}

	return wordPatternMatchHelper(pattern, str, keys, vals)
}

func wordPatternMatchHelper(pattern, str string, keys, vals map[string]string) bool {
	if len(pattern) == 0 {
		return len(str) == 0
	}

	for i := 0; i < len(str); i++ {
		tmp := str[:i+1]
		if keys[string(pattern[0])] == "" && vals[tmp] == "" {

			keys[string(pattern[0])] = tmp
			vals[tmp] = string(pattern[0])

			if wordPatternMatchHelper(pattern[1:], str[i+1:], keys, vals) {
				return true
			}
			delete(keys, string(pattern[0]))
			delete(vals, tmp)
		} else if keys[string(pattern[0])] == tmp && vals[tmp] == string(pattern[0]) {
			if wordPatternMatchHelper(pattern[1:], str[i+1:], keys, vals) {
				return true
			}
		}
	}

	return false
}
