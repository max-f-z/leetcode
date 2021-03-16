package main

func medianSlidingWindow(nums []int, k int) []float64 {
	mf := Constructor295()

	res := []float64{}

	for i := 0; i < k; i++ {
		mf.AddNum(nums[i])
	}

	// k* log(k)

	res = append(res, mf.FindMedian())

	// n * (k + log(k))
	for i := 1; i < len(nums)-k+1; i++ {
		mf.Remove(nums[i-1])
		mf.AddNum(nums[i+k-1])
		res = append(res, mf.FindMedian())
	}

	return res
}
