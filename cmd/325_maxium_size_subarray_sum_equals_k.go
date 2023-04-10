package main

//lint:ignore U1000 unused
func maxSubArrayLen(nums []int, k int) int {
	record := map[int]int{}
	record[0] = -1
	sum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if _, ok := record[sum]; !ok {
			record[sum] = i
		}

		if tmp, ok := record[sum-k]; ok {
			if max < (i - tmp) {
				max = i - tmp
			}
		}
	}
	return max
}
