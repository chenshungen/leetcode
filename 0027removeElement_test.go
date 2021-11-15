package leetcode

import "testing"

func removeElement(nums []int, val int) int {
	nl := len(nums)
	if nl <= 0 {
		return 0
	}

	i, j := 0, nl-1
	for i = 0; i <= j && i < nl && j >= 0; {
		if nums[i] == val {
			if nums[j] == val {
				j--
				continue
			}
			nums[j], nums[i] = nums[i], nums[j]
			j--
		} else {
			i++
		}
	}

	nums = nums[:i]

	return i
}

func TestRemoveElement(t *testing.T) {
	t.Log(removeElement([]int{3, 2, 3, 2, 3, 3, 3, 2, 2, 3}, 3))
}
