package leetcode

import (
	"sort"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	l3 := l1 + l2
	mergeNums := make([]int, l3)
	copy(mergeNums[:l1], nums1)
	copy(mergeNums[l1:l1+l2], nums2)
	sort.Sort(sort.IntSlice(mergeNums))
	if l3%2 == 0 {
		mid := float64(mergeNums[(l3/2)-1]+mergeNums[l3/2]) / 2
		return mid
	} else {
		return float64(mergeNums[l3/2])
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	t.Log(findMedianSortedArrays(nums1, nums2))
}
