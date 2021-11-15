package leetcode

import "testing"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	res := make([]int, 0, 10)
	for x != 0 {
		res = append(res, x%10)
		x = x / 10
	}

	l, r := 0, len(res)-1
	for l < r {
		if res[l] == res[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome(1111))
}
