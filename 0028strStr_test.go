package leetcode

import (
	"strings"
	"testing"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	var i int
	for i = 0; i < len(haystack); {
		if haystack[i] != needle[0] {
			i++
			continue
		}

		// 第一个字符相同后开始迭代后续的是否全部相同
		allSame := true
		for j := 1; j < len(needle); j++ {
			if (i+j >= len(haystack)) || (i+j < len(haystack) && haystack[i+j] != needle[j]) {
				allSame = false
				break
			}
		}

		if allSame {
			return i
		} else {
			i++
		}
	}

	return -1
}

func TestStrStr(t *testing.T) {
	t.Log(strings.Index("aaaaaaaaaaai", "i"))
	t.Log(strStr("aaaaaaaaaaai", "i"))
}
