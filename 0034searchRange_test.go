package leetcode

import "testing"

func searchRange(nums []int, target int) []int {
	//长度小于等于0 ，那直接嗝屁
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	// 二分查找
	index := bSearch(nums, target)
	// 找不到 index 输出负一
	if index == -1 {
		return []int{-1, -1}
	}
	// 前后移动
	min, max := index, index
	for {
		//注意不要下标溢出
		if min-1 >= 0 && nums[min-1] == nums[index] {
			min--
			continue
		}
		break
	}
	for {
		//注意不要下标溢出
		if max+1 <= len(nums)-1 && nums[max+1] == nums[index] {
			max++
			continue
		}
		break
	}
	return []int{min, max}
}

func TestSearchRange(t *testing.T) {
	t.Log(searchRange([]int{1, 3}, 1))
	//t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))
}
