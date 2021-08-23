package main

import (
	"fmt"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {

	l := len(s)
	if l <= 0 {
		return s, 0
	}

	return s[:l-1], s[l-1]
}

func (s stack) Top() int {
	l := len(s)
	if l <= 0 {
		return 0
	}

	return s[l-1]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//	https://leetcode.com/problems/longest-valid-parentheses/submissions/
func longestValidParentheses(s string) int {
	maxans := 0
	stk := make(stack, 0, len(s))

	stk = stk.Push(-1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = stk.Push(i)
		} else {
			stk, _ = stk.Pop()
			if len(stk) == 0 {
				stk = stk.Push(i)
			} else {
				maxans = max(maxans, i-stk.Top())
			}
		}
	}
	return maxans
}

func main() {
	str1 := "(())"
	fmt.Println(longestValidParentheses(str1))
}
