package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int

	for i < len(nums)-1 && nums[i] <= nums[i+1] {
		i++
	}

	if i+1 <= len(nums)-1 {
		return nums[i+1]
	}

	return nums[0]
}

func testFindMin153() {
	nums := []int{1, 3, 5}
	fmt.Println(findMin(nums))
}
