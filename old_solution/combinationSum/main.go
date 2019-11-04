package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/combination-sum/submissions/
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	set := []int{}
	explore(candidates, target, 0, &result, set)

	return result
}

func explore(candidates []int, target int, pos int, result *[][]int, set []int) {
	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(set))
		copy(c, set)
		*result = append(*result, c)
		return
	}

	for i := pos; i < len(candidates); i++ {
		set = append(set, candidates[i])
		// 逐个元素递减(小于等于时返回表示找到了一组)
		explore(candidates, target-candidates[i], i, result, set)
		// 去掉最后一个元素 重新递归查找
		set = set[:len(set)-1]
	}
}

func combinationSum2(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	combinationSum2Helper(nums, nil, target, 0, 0, &result)
	return result
}

func combinationSum2Helper(nums, combo []int, target, sum, startIndex int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, combo...))
		return
	}
	for i := startIndex; i < len(nums) && (sum+nums[i]) <= target; i++ {
		if i != startIndex && nums[i] == nums[i-1] {
			continue
		}
		combinationSum2Helper(nums, append(combo, nums[i]), target, sum+nums[i], i+1, result)
	}
}

func main() {

	nums := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(nums, 8))

}
