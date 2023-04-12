package main

//lint:ignore U1000 unused
func findSmallestInteger(nums []int, value int) int {
	n := len(nums)
	ref := map[int]int{}

	for i := 0; i < n; i++ {
		local := (nums[i]%value + value) % value

		ref[local]++
	}
	for i := 0; i < n; i++ {
		local := i % value
		ref[local]--
		if ref[local] < 0 {
			return i
		}
	}
	return n
}
