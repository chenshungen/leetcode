package leetcode

import "testing"

func searchInsert(nums []int, target int) int {
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

	return l
}

func TestSearchInsert(t *testing.T) {
	t.Log(searchInsert([]int{1, 3, 5, 6}, 0))
}
