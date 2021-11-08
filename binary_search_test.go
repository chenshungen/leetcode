package leetcode

import "testing"

func bSearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	return -1
}

func TestBSearch(t *testing.T) {
	t.Log(bSearch([]int{1, 3, 5, 6}, 10))
	//t.Log(bSearch([]int{1, 2, 5, 10, 19, 21, 29}, 29))
	// t.Log(bSearch([]int{-1, -2, 0, 5, 10, 19, 21, 29}, 10))
	// t.Log(bSearch([]int{}, 10))
}
