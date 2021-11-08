package leetcode

import (
	"sort"
	"testing"
)

// 排序+双指针
func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	// 先排序将数据从小到大排列  [-4,-1,-1,0,1,2]
	sort.Ints(nums)
	numsLength := len(nums)

	var res [][]int
	for k := 0; k < numsLength-2; k++ {

		if nums[k] > 0 {
			break
		}

		// 跳过重复的数
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		left := k + 1
		right := numsLength - 1

		for left < right {

			sum := nums[k] + nums[left] + nums[right]
			if sum > 0 {
				for right = right - 1; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else if sum < 0 {
				for left = left + 1; left < right && nums[left] == nums[left-1]; left++ {
				}
			} else {
				res = append(res, []int{nums[k], nums[left], nums[right]})
				for left = left + 1; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right = right - 1; left < right && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}

	return res
}

func TestThreeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
