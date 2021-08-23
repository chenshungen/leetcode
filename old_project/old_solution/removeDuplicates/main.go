package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(nums []int) int {
	res := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if i >= 1 && nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			res--
		}
	}

	return res
}

func removeElement(nums []int, val int) int {
	res := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			res--
		}
	}

	return res
}

func removeDuplicates2(nums []int) int {
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return len(nums)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pCurr := head

	for pCurr != nil {
		for pCurr.Next != nil && pCurr.Val == pCurr.Next.Val {
			pCurr.Next = pCurr.Next.Next
		}
		pCurr = pCurr.Next
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	// 长度 <=1 的 list ，可以直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 要么 head 重复了，那就删除 head
	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates2(head.Next)
	}

	// 要么 head 不重复，递归处理 head 后面的节点
	head.Next = deleteDuplicates2(head.Next)
	return head
}

func printList(pHead *ListNode) {
	for pHead != nil {
		if pHead.Next == nil {
			fmt.Printf("%v", pHead.Val)
		} else {
			fmt.Printf("%v->", pHead.Val)
		}
		pHead = pHead.Next
	}
	fmt.Println()
}

func main() {

	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3}

	// //len := removeDuplicates(nums)

	// //len := removeElement(nums, 1)

	// len := removeDuplicates2(nums)

	// for i := 0; i < len; i++ {
	// 	fmt.Println(nums[i])
	// }

	head := &ListNode{Val: 1, Next: nil}
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 3, Next: nil}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	printList(deleteDuplicates(head))
}
