package main

import "fmt"

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}

	if needleLen > haystackLen {
		return -1
	}

	index := 0
	isMatch := true
	for index < haystackLen {
		if haystack[index] != needle[0] {
			index++
			continue
		}

		nextFirstMatch := -1
		for i := 0; i < needleLen; i++ {
			if index+i >= haystackLen {
				return -1
			}
			if haystack[index+i] != needle[i] {
				isMatch = false
			}
			if i > 0 && nextFirstMatch == -1 && haystack[index+i] == needle[0] {
				nextFirstMatch = index + i
			}
		}

		if !isMatch {
			isMatch = true
			if nextFirstMatch != -1 {
				index = nextFirstMatch
			} else {
				index++
			}
		} else {
			return index
		}
	}

	return -1
}

func main() {
	mainStr := "aabaaabaaac"
	subStr := "aabaaac"
	fmt.Println(strStr(mainStr, subStr))
}
