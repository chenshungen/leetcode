package main

import (
	"fmt"
	"sort"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	if len(intervals) == 2 {
		// 有重叠需要合并
		if intervals[0][1] >= intervals[1][0] {
			intervals[0][0] = min(intervals[0][0], intervals[1][0])
			intervals[0][1] = max(intervals[0][1], intervals[1][1])
			intervals = append(intervals[:1], intervals[2:]...)
		}
		return intervals
	}

	for i := 1; i < len(intervals); {
		if intervals[i-1][1] >= intervals[i][0] {
			intervals[i-1][0] = min(intervals[i-1][0], intervals[i][0])
			intervals[i-1][1] = max(intervals[i-1][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			i++
		}
	}

	return intervals
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := 0, 0
	for p1 < m+p2 && p2 < n {
		for p1 < m+p2 && nums1[p1] <= nums2[p2] {
			p1++
		}
		copy(nums1[p1+1:], nums1[p1:])
		nums1[p1] = nums2[p2]
		p2++
	}

	if p2 < n {
		p1 = m + p2
		for p2 < n {
			nums1[p1] = nums2[p2]
			p1++
			p2++
		}
	}
}

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// intervals := [][]int{{1, 4}, {0, 0}}
	// intervals = merge(intervals)
	//fmt.Println(merge(intervals))
	//fmt.Println(intervals)

	nums1 := []int{4, 0, 0, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 5, 6}
	merge2(nums1, 1, nums2, 5)
	fmt.Println(nums1)
}
