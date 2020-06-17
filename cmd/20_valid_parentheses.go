package main

func isValid(s string) bool {
	local := ""
	dict := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for _, v := range s {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			local += dict[string(v)]
		} else {
			if local == "" {
				return false
			}
			if string(local[len(local)-1]) == string(v) {
				local = local[:len(local)-1]
			} else {
				return false
			}
		}
	}

	return local == ""
}
