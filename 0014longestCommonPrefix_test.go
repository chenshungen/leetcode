package leetcode

import (
	"testing"
)

func longestCommonPrefix(strs []string) string {

	minLen := 200
	for _, str := range strs {
		sl := len(str)
		if sl <= minLen {
			minLen = sl
		}
	}

	var res string
	for i := 0; i < minLen; i++ {
		flag := strs[0][i]
		allSame := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != flag {
				allSame = false
				break
			}
		}

		if allSame {
			res += string(flag)
		} else {
			break
		}
	}

	return res
}

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
