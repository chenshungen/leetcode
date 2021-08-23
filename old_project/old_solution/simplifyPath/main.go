package main

import "fmt"

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

func simplifyPath(path string) string {
	if len(path) <= 1 {
		return "/"
	}
	low := 1
	high := len(path) - 1
	stk := make(stack, 0, len(path))
	stk = stk.Push("/")

	// 跳过开头所有/
	for low <= high && path[low] == '/' {
		low++
	}

	// 跳过末尾所有
	for high >= 0 && path[high] == '/' {
		high--
	}

	var word string
	for low <= high {
		if path[low] == '/' {
			for low <= high && path[low] == '/' {
				low++
			}
			if word == ".." {
				if len(stk) > 1 {
					stk, _ = stk.Pop()
				}
			} else if word != "." {
				stk = stk.Push(word)
			}
			word = ""
		} else {
			word += string(path[low])
			low++
		}

	}

	if word != "" && word != "." && word != ".." {
		stk = stk.Push(word)
	}

	if word == ".." {
		if len(stk) > 1 {
			stk, _ = stk.Pop()
		}
	}

	var res string
	for i := 0; i < len(stk); i++ {
		str := stk[i]
		if i == 0 || i == len(stk)-1 {
			res += str
		} else {
			res += str + "/"
		}
	}

	return res
}

func main() {
	// fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	// fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	//fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath("///"))
}
