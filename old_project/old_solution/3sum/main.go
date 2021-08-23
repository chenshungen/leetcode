package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var (
		a   int
		b   int
		c   int
		res [][]int
	)
	res = make([][]int, 0, 0)
	numsLen := len(nums)
	if numsLen < 3 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < numsLen-2; i++ {

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		a = nums[i]
		low := i + 1
		high := numsLen - 1
		for low < high {
			b = nums[low]
			c = nums[high]
			if a+b+c == 0 {
				element := make([]int, 0, 3)
				element = append(element, a, b, c)
				res = append(res, element)
				for low < numsLen-1 && nums[low] == nums[low+1] {
					low++
				}
				for high > 0 && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			} else if a+b+c > 0 {
				for high > 0 && nums[high] == nums[high-1] {
					high--
				}
				high--
			} else {
				for low < numsLen-1 && nums[low] == nums[low+1] {
					low++
				}
				low++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	res := threeSum(nums)
	fmt.Println(res)
}
