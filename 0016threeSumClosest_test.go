package leetcode

import (
	"math"
	"sort"
	"testing"
)

//  排序 + 对撞双指针 + 去重优化
// 时间复杂度：O(n^2)  空间复杂度：O(1)

func threeSumClosest(nums []int, target int) int {
	res := 0
	sort.Ints(nums)
	dist := math.MaxInt32 // 记录当前三数之和与目标值的距离，距离越小越接近目标值

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // 对第一个元素作去重优化
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum > target {
				if sum-target < dist {
					dist = sum - target
					res = sum
				}
				for k > j && nums[k-1] == nums[k] { // 对第二个元素作去重优化
					k--
				}
				k--
			} else {
				if target-sum < dist {
					dist = target - sum
					res = sum
				}
				for j < k && nums[j+1] == nums[j] { // 对第三个元素作去重优化
					j++
				}
				j++
			}
		}
	}

	return res
}

func TestThreeSumClosest(t *testing.T) {
	t.Log(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
