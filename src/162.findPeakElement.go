package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if tmp*diff >= 0 {
			tmp += diff
			continue
		}

		return i
	}

	return -1
}

func testFindPeakElement162() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}
