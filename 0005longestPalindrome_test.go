package leetcode

import (
	"testing"
)

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	r := make([][]bool, l)
	for i := 0; i < l; i++ {
		r[i] = make([]bool, l)
	}

	ml, si := 1, 0
	// 字串长度 sl = j - i + 1
	for sl := 2; sl <= l; sl++ {
		// 左边界 i
		for i := 0; i < l; i++ {
			// 右边界 j - i + 1 = sl
			j := i + sl - 1
			if j >= l {
				break
			}
			if s[i] != s[j] {
				r[i][j] = false
			} else {
				if sl <= 3 {
					r[i][j] = true
				} else {
					r[i][j] = r[i+1][j-1]
				}
			}

			if r[i][j] && sl > ml {
				ml = sl
				si = i
			}
		}
	}
	return s[si : si+ml]
}

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("babad"))
}
