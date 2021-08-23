package main

import (
	"fmt"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {

	l := len(s)
	if l <= 0 {
		return s, ""
	}

	return s[:l-1], s[l-1]
}

// https://leetcode.com/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	strLen := len(s)
	if strLen <= 0 {
		return ""
	}

	low := 0
	high := strLen - 1
	for low <= high && s[low] == ' ' {
		low++
	}
	for low <= high && s[high] == ' ' {
		high--
	}

	var word string
	stk := make(stack, 0, len(s))
	for low <= high {
		if s[low] == ' ' {
			for low <= high && s[low] == ' ' {
				low++
			}
			stk = stk.Push(word)
			word = ""
		} else {
			word += string(s[low])
			low++
		}
	}
	if word != "" {
		stk = stk.Push(word)
	}

	var res string
	for len(stk) > 0 {
		stk, word = stk.Pop()
		if len(stk) > 0 {
			res += word + " "
		} else {
			res += word
		}
	}

	return res
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
