package main

import (
	"fmt"
	"math"
)

func main() {

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(12 << uint(2))

	fmt.Println(math.Pow(144, 1/float64(64)))

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
