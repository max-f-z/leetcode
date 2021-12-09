package main

func canReach(arr []int, start int) bool {
	queue := []int{start}

	length := len(arr)

	for len(queue) != 0 {
		newQ := []int{}

		for _, v := range queue {
			if arr[v] == -1 {
				continue
			}

			value := arr[v]

			if v+value < length && v+value >= 0 && arr[v+value] != -1 {
				if arr[v+value] == 0 {
					return true
				}
				newQ = append(newQ, v+value)
			}

			if v-value < length && v-value >= 0 && arr[v-value] != -1 {
				if arr[v-value] == 0 {
					return true
				}
				newQ = append(newQ, v-value)
			}

			arr[v] = -1
		}

		queue = newQ
	}

	return false
}
