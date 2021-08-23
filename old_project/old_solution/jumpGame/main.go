package main

import "fmt"

func canJump(nums []int) bool {

	for i := len(nums) - 2; i >= 0; i-- {
		// 找到数值为 0 的元素
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if i-j < nums[j] {
				// 在 j 号位置上，可以跨过 0 元素
				i = j
				break
			}
		}

		if j == -1 {
			// 在 0 元素之前，没有位置可以跨过 0
			return false
		}
	}

	return true
}

func jump(nums []int) int {
	i, count, end := 0, 0, len(nums)-1

	var nextI, maxNextI, maxI int
	for i < end {
		if i+nums[i] >= end {
			return count + 1
		}

		nextI, maxNextI = i+1, i+nums[i]
		for nextI <= maxNextI {
			if nextI+nums[nextI] > maxI {
				maxI, i = nextI+nums[nextI], nextI
			}

			nextI++
		}

		count++
	}

	return count
}

func main() {
	nums := []int{2, 3, 1, 0, 4}
	//fmt.Println(canJump(nums))
	fmt.Println(jump(nums))
}
