package main

//lint:ignore U1000 unused
func findStrobogrammatic(n int) []string {

	return findStrobogrammaticHelper(n, n)
}

func findStrobogrammaticHelper(n, m int) []string {
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	if n == 0 {
		return []string{""}
	}

	preList := findStrobogrammaticHelper(n-2, m)

	curList := make([]string, 0)

	for _, v := range preList {
		if n != m {
			curList = append(curList, "0"+v+"0")
		}
		curList = append(curList, "1"+v+"1")
		curList = append(curList, "6"+v+"9")
		curList = append(curList, "8"+v+"8")
		curList = append(curList, "9"+v+"6")
	}
	return curList
}
