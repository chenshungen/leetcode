package main

import "fmt"

//https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	low := 0
	mid := 0
	high := len(nums) - 1
	pivot := 0
	for low < high {
		mid = (high + low) >> 1
		if low > 1 && nums[mid] < nums[mid-1] {
			pivot = mid
			break
		} else {
			low++
		}
	}

	low = 0
	mid = 0
	high = pivot - 1
	for low < high {
		mid = (high + low) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else if nums[mid] < target {
			low = mid
		}
	}

	low = pivot
	mid = 0
	high = len(nums) - 1
	for low < high {
		mid = (high + low) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else if nums[mid] < target {
			low = mid
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	k := 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}

	i, j := 0, length-1
	for i <= j {
		m := (i + j) / 2
		med := (m + k) % length

		switch {
		case nums[med] < target:
			i = m + 1
		case target < nums[med]:
			j = m - 1
		default:
			return med
		}
	}

	return -1
}

func search3(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	k := 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}

	i, j := 0, length-1
	for i <= j {
		m := (i + j) / 2
		med := (m + k) % length

		switch {
		case nums[med] < target:
			i = m + 1
		case target < nums[med]:
			j = m - 1
		default:
			return true
		}
	}

	return false
}

func searchRange(nums []int, target int) []int {
	res := make([]int, 0, 2)
	if len(nums) == 0 {
		res = append(res, -1, -1)
		return res
	}
	low := 0
	high := len(nums) - 1

	var find bool
	var mid int
	for low <= high {
		mid = (high + low) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] == target {
			find = true
			break
		}
	}

	if !find {
		res = append(res, -1, -1)
	} else {
		tmp := mid
		for tmp > 0 && nums[tmp] == nums[tmp-1] {
			tmp--
		}
		res = append(res, tmp)

		tmp = mid
		for tmp < len(nums)-1 && nums[tmp] == nums[tmp+1] {
			tmp++
		}
		res = append(res, tmp)
	}

	return res
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] {
		return 0
	}

	index := 0
	for index < len(nums) {
		if nums[index] == target {
			return index
		}
		if index < len(nums)-1 &&
			nums[index] < target && nums[index+1] > target {
			return index + 1
		}
		index++
	}
	return index
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 0))
}
