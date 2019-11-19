package main

import "math"

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

func maxProduct(a []int) int {
	cur, neg, max := 1, 1, a[0]

	for i := 0; i < len(a); i++ {

		switch {
		case a[i] > 0:
			cur, neg = a[i]*cur, a[i]*neg
		case a[i] < 0:
			cur, neg = a[i]*neg, a[i]*cur
		default:
			cur, neg = 0, 1
		}

		if max < cur {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}
