package leetcode

import (
	"testing"
)

func twoSum(nums []int, target int) []int {

	numsLen := len(nums)

	if numsLen < 2 {
		return []int{}
	}

	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func TestTowSum(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{3, 2, 4}, 6))
	t.Log(twoSum([]int{3, 3}, 6))
}
