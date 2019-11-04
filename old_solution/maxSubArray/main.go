package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	temp := 0
	sumMax := math.MinInt32
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		if temp > sumMax {
			sumMax = temp
		}
		if temp < 0 {
			temp = 0
		}
	}

	return sumMax
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
