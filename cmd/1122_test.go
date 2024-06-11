package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1122(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}

	ans := relativeSortArray(arr1, arr2)
	assert.DeepEqual(t, ans, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19})

	arr1 = []int{28, 6, 22, 8, 44, 17}
	arr2 = []int{22, 28, 8, 6}

	ans = relativeSortArray(arr1, arr2)
	assert.DeepEqual(t, ans, []int{22, 28, 8, 6, 17, 44})
}
