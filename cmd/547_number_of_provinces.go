package main

func findCircleNum(isConnected [][]int) int {
	p := 0

	n := len(isConnected)

	provinces := make([]int, n)

	for i := 0; i < n; i++ {
		if provinces[i] == 0 {
			p++
			provinces[i] = p
			provinces = findCircleNumHelper(p, i, provinces, isConnected)
		}
	}

	return p
}

func findCircleNumHelper(p int, city int, provinces []int, isConnected [][]int) []int {
	for i := 0; i < len(isConnected); i++ {
		if (isConnected[city][i] == 1 || isConnected[i][city] == 1) && provinces[i] == 0 && i != 0 {
			provinces[i] = p
			provinces = findCircleNumHelper(p, i, provinces, isConnected)
		}
	}
	return provinces
}
