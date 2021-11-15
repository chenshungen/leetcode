package leetcode

import "testing"

type Stack struct {
	i    int
	data []byte
}

func newStack() *Stack {
	return &Stack{
		i:    0,
		data: make([]byte, 0, 8),
	}
}

func (s *Stack) Push(k byte) {
	s.data = append(s.data, k)
	s.i++
}

func (s *Stack) Pop() (ret byte) {
	ret = s.data[s.i-1]
	s.data = s.data[:s.i-1]
	s.i--
	return
}

func (s *Stack) Top() (ret byte) {
	ret = s.data[s.i-1]
	return
}

func (s *Stack) Empty() bool {
	return s.i == 0
}

func isMatch(a, b byte) bool {
	return (a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}')
}

func isValid(s string) bool {

	stack := newStack()
	for i := range s {
		if stack.Empty() {
			stack.Push(s[i])
		} else {
			top := stack.Top()
			if isMatch(top, s[i]) {
				stack.Pop()
			} else {
				stack.Push(s[i])
			}
		}
	}
	return stack.Empty()
}

func TestIsValid(t *testing.T) {
	t.Log(isValid("()[]{}"))
}
