package leetcode

import (
	"testing"
)

func maxArea(height []int) int {
	minFunc := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	maxFunc := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var max int
	l, r := 0, len(height)-1
	for l < r {
		max = maxFunc(max, (r-l)*minFunc(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
