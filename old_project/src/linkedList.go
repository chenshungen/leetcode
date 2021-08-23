package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// LinkedList ...
type LinkedList struct {
	head   *ListNode
	length uint
}

// NewLinkedList ...
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// NewListNode ...
func NewListNode(v interface{}) *ListNode {
	return &ListNode{v, nil}
}

// InsertAfter ...
func (list *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.Next
	p.Next = newNode
	newNode.Next = oldNext
	list.length++
	return true
}

// InsertBefore 在某节点前插入
func (list *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == list.head {
		return false
	}
	cur := list.head.Next
	pre := list.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur
	list.length++
	return true
}

// InsertToHead 在链表头插入节点
func (list *LinkedList) InsertToHead(v interface{}) bool {
	return list.InsertAfter(list.head, v)
}

// InsertToTail 在链表尾插入节点
func (list *LinkedList) InsertToTail(v interface{}) bool {
	cur := list.head
	for nil != cur.Next {
		cur = cur.Next
	}
	return list.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找节点
func (list *LinkedList) FindByIndex(idx uint) *ListNode {
	if idx >= list.length {
		return nil
	}
	cur := list.head.Next
	var i uint
	for i < idx {
		cur = cur.Next
		i++
	}
	return cur
}

// DeleteNode 删除节点
func (list *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := list.head.Next
	pre := list.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if nil == cur {
		return false
	}
	pre.Next = p.Next
	p = nil
	list.length--
	return true
}

//Print 打印链表
func (list *LinkedList) Print() {
	cur := list.head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Val)
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
