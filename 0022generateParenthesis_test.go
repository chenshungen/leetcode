package leetcode

import (
	"testing"
)

func generateParenthesis(n int) []string {
	set := make(map[string]bool)
	res := make([]string, 0)

	var DFS func(depth int, s string)
	DFS = func(depth int, s string) {
		if set[s] == true {
			return
		}
		set[s] = true
		if depth == n {
			res = append(res, s)
			return
		}

		str := "()" + s
		DFS(depth+1, str)
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				str = s[:i+1] + "()" + s[i+1:]
				DFS(depth+1, str)
			}
		}
	}

	DFS(0, "")

	return res
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
}
