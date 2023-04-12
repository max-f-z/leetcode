package main

import "strings"

func simplifyPath(path string) string {
	directories := strings.Split(path, "/")

	stack := make([]string, 0)
	for _, directory := range directories {
		if directory == "" || directory == "." {
			// empty or current directory
			continue
		}

		if directory == ".." {
			// in order to go back to parent directory
			if 0 < len(stack) {
				// remove last index in stack
				stack = append(stack[:len(stack)-1], stack[len(stack)-1+1:]...)
			}
		} else {
			stack = append(stack, directory)
		}
	}

	// root "/" + child path
	return "/" + strings.Join(stack, "/")
}
