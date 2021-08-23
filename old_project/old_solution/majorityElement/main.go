package main

import "fmt"

// https://leetcode.com/problems/majority-element/
func majorityElement(nums []int) int {
	var (
		tmpMap   map[int]int
		max      int
		majority int
	)
	tmpMap = make(map[int]int)
	for _, num := range nums {
		if _, ok := tmpMap[num]; ok {
			tmpMap[num]++
		} else {
			tmpMap[num] = 0
		}
	}
	max = -1
	for key, value := range tmpMap {
		if value > max {
			max = value
			majority = key
		}
	}

	return majority
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	fmt.Println(res)
}
