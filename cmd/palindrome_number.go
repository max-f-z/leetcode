package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	save := x
	for {
		if save == 0 {
			break
		}
		reverse = reverse*10 + (save % 10)
		save = save / 10
	}

	return reverse == x
}
