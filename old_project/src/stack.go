package main

import "fmt"

// Stack 栈接口
type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush()
}

// ArrayStack 基于数组实现的栈
type ArrayStack struct {
	data []interface{}
	top  int
}

// NewArrayStack ...
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// IsEmpty 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return true
	}
	return false
}

// Push 压栈
func (stack *ArrayStack) Push(v interface{}) {
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top++
	}
	if stack.top > len(stack.data)-1 {
		stack.data = append(stack.data, v)
	} else {
		stack.data[stack.top] = v
	}
}

// Pop 出栈
func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	v := stack.data[stack.top]
	stack.top--
	return v
}

// Top 栈顶元素
func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

// Flush 刷新栈数据
func (stack *ArrayStack) Flush() {
	stack.top = -1
}

// Print 打印栈数据
func (stack *ArrayStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := stack.top; i >= 0; i-- {
			fmt.Println(stack.data[i])
		}
	}
}
