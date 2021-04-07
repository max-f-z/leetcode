package main

import "strings"

func findDuplicate(paths []string) [][]string {
	cols := map[string][]string{}

	for _, path := range paths {
		strs := strings.Split(path, " ")
		prefix := strs[0]

		for i := 1; i < len(strs); i++ {
			contentStrs := strings.Split(strs[i], "(")
			fileName := contentStrs[0]
			fileContent := contentStrs[1][:len(contentStrs[1])-1]
			sb := strings.Builder{}
			sb.WriteString(prefix)
			sb.WriteByte('/')
			sb.WriteString(fileName)
			if cols[fileContent] == nil {

				cols[fileContent] = []string{sb.String()}
			} else {
				cols[fileContent] = append(cols[fileContent], sb.String())
			}
		}
	}

	res := [][]string{}

	for k := range cols {
		if len(cols[k]) > 1 {
			res = append(res, cols[k])
		}
	}

	return res
}
