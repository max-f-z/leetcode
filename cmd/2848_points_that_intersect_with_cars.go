package main

func numberOfPoints(nums [][]int) int {
	place := make([]int, 102)

	for i := 0; i < len(nums); i++ {
		place[nums[i][0]] += 1
		place[nums[i][1]+1] -= 1
	}

	res, cars := 0, 0

	for i := 0; i < 101; i++ {
		cars += place[i]

		if cars > 0 {
			res += 1
		}
	}

	return res
}
