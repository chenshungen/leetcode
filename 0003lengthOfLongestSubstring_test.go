package leetcode

import "testing"

// 滑动窗口法
func lengthOfLongestSubstring(s string) int {
	maxFunc := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	lookup := make(map[byte]struct{})
	l, r, maxLen := 0, 0, 0
	for l < len(s) && r < len(s) {
		if _, ok := lookup[s[r]]; !ok {
			lookup[s[r]] = struct{}{}
			maxLen = maxFunc(maxLen, (r - l + 1))
			r++
		} else {
			delete(lookup, s[l])
			l++
		}
	}

	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcabccccaa"))
}
