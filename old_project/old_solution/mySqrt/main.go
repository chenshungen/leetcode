package main

import "fmt"

// https://leetcode.com/problems/sqrtx/
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	var (
		low  = 0
		high = x
		mid  = 0
	)

	for low < high {
		mid = (high + low) >> 1
		if (mid-1)*(mid-1) < x && (mid+1)*(mid+1) > x {
			if mid*mid <= x {
				return mid
			}
			return mid - 1
		} else if mid*mid > x {
			high = mid
		} else if mid*mid < x {
			low = mid
		}
	}
	return mid
}

func main() {
	fmt.Println(mySqrt(6))
}
