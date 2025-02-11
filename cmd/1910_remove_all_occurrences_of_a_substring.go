package main

func removeOccurrences(s string, part string) string {
	stack := []rune{}

	for _, v := range s {
		stack = append(stack, v)

		if len(stack) < len(part) {
			continue
		}

		tail := string(stack[len(stack)-len(part):])

		if tail == part {
			stack = stack[:len(stack)-len(part)]
		}
	}

	return string(stack)
}
