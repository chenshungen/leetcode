package main

import (
	"fmt"
)

// https://leetcode.com/problems/first-missing-positive/
func firstMissingPositive(nums []int) int {
	// 整理 nums ，让 nums[k] == k+1，只要 k+1 存在于 nums 中
	for i := 0; i < len(nums); i++ {

		// fmt.Println(i)

		for 0 <= nums[i]-1 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			// 当 for 的判断语句成立时，
			// nums[i]-1 就是 k ，nums[i] 的值是 k+1
			// nums[i] != nums[nums[i]-1] 即是 k+1 != nums[k] ，这说明
			// 1. k+1 存在与 nums 中，
			// 2. k+1 还没有在他该在的 nums[k] 中
			// 通过互换，让 k+1 到 nums[k] 中去
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

			// fmt.Println(nums)

			// 使用 for 而不是 if 是因为
			// nums[i] 中的新值，有可能是另一个 k+1 ，需要再让其归为
			// 如果使用 if ，而这个新的 k+1 又只有一个的话，
			// 这个新的 k+1 不会被处理到，不会被放在 nums[k] 中
		}
	}
	// 循环结束后，所有 1<=k+1<=len(nums) 且 k+1 存在于nums中，都会被存放于 nums[k] 中

	// 整理后，第一个不存在的 k+1 就是答案
	for k := range nums {
		if nums[k] != k+1 {
			return k + 1
		}
	}
	return len(nums) + 1
}

// https://leetcode.com/problems/first-missing-positive/
func firstMissingPositive2(nums []int) int {
	arrayMap := make(map[int]bool)
	var max int

	// Store all positive numbers in a map and find out the maximum number within this array
	for _, v := range nums {
		// Choose the maximum number
		if v > max {
			max = v
		}

		// Filter out all negative numbers
		if v > 0 {
			arrayMap[v] = true
		}
	}

	// Iterate and seek out the first missing positive based on the map initialized above
	for i := 1; i <= max; i++ {
		if _, ok := arrayMap[i]; !ok {
			return i
		}
	}

	// The first missing positive would be the integer right after the maximum number in this array,
	// if we can't hit it before reaching the maximum number.
	return max + 1
}

func main() {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	res := firstMissingPositive2(nums)
	fmt.Println(res)
}
