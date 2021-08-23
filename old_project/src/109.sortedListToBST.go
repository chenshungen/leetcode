package main

func testSortedListToBST109() {
	list := NewLinkedList()
	list.InsertToTail(-10)
	list.InsertToTail(-3)
	list.InsertToTail(0)
	list.InsertToTail(5)
	list.InsertToTail(9)
	list.Print()

	bTree := NewBinaryTree(0)
	bTree.Root = sortedListToBST(list.head.Next)
	bTree.InOrderTraverse()
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := make([]int, 0, 0)
	pCurr := head
	for pCurr != nil {
		nums = append(nums, pCurr.Val.(int))
		pCurr = pCurr.Next
	}
	return sortedArrayToBSTHelper(nums)
}

func sortedArrayToBSTHelper(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums) - 1) / 2
	root := NewTreeNode(nums[mid])
	if len(nums) == 1 {
		return root
	}
	root.Left = sortedArrayToBSTHelper(nums[:mid])
	root.Right = sortedArrayToBSTHelper(nums[mid+1:])
	return root
}
