package main

import (
	"datastructures"
	"fmt"
	"strconv"
	"strings"
)

func isSameTree(p *datastructures.TreeNode, q *datastructures.TreeNode) bool {
	if p == q && p == nil {
		return true
	}
	pIter, qIter := p, q
	for pIter != nil && qIter != nil {
		if pIter.Val == qIter.Val {
			return isSameTree(pIter.Left, qIter.Left) &&
				isSameTree(pIter.Right, qIter.Right)
		} else {
			return false
		}
	}
	return false
}

func levelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0, 1)
	queue1 := make([]*datastructures.TreeNode, 0, 2)
	queue1 = append(queue1, root)

	for len(queue1) > 0 {
		ret := make([]int, 0, 0)
		queue2 := make([]*datastructures.TreeNode, 0, 2)
		for _, node := range queue1 {
			ret = append(ret, node.Val.(int))
			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}

		queue1 = queue2
		res = append(res, ret)
	}

	// for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
	// 	res[left], res[right] = res[right], res[left]
	// }

	return res
}

func maxDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func minDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	} else {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildTreeFromPreAndInOrder(preorder []int, inorder []int) {
	binTree := datastructures.NewBinaryTree(preorder[0])
	binTree.Root = buildTree(preorder, inorder)
	binTree.PreOrderTraverse()
	binTree.InOrderTraverse()
}

func buildTree(preorder []int, inorder []int) *datastructures.TreeNode {
	preLen, inLen := len(preorder), len(inorder)
	if preLen == 0 || inLen == 0 {
		return nil
	}
	root := datastructures.NewTreeNode(preorder[0])
	// 中序遍历序列根节点的位置
	var in int
	for inorder[in] != preorder[0] {
		in++
	}
	inorderLeft := make([]int, 0, inLen)
	inorderRight := make([]int, 0, inLen)
	if in != 0 {
		inorderLeft = inorder[:in]
	}
	if in != inLen-1 {
		inorderRight = inorder[in+1:]
	}
	root.Left = buildTree(preorder[1:in+1], inorderLeft)
	root.Right = buildTree(preorder[in+1:], inorderRight)
	return root
}

func sortedArrayToBST(nums []int) *datastructures.TreeNode {
	mid := (len(nums) - 1) / 2
	root := datastructures.NewTreeNode(nums[mid])
	if len(nums) == 1 {
		return root
	}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func binaryTreePaths(root *datastructures.TreeNode) [][]int {
	res := make([][]int, 0, 0)
	s := []string{}
	if root == nil {
		return [][]int{}
	}
	helper(root, strconv.Itoa(root.Val.(int)), &s)

	for _, path := range s {
		ps := strings.Split(path, "->")
		psi := make([]int, 0, len(ps))
		for _, v := range ps {
			intV, _ := strconv.ParseInt(v, 10, 32)
			psi = append(psi, int(intV))
		}
		res = append(res, psi)
	}

	return res
}

func helper(node *datastructures.TreeNode, path string, s *[]string) {
	if node.Left == nil && node.Right == nil {
		*s = append(*s, path)
	}
	if node.Left != nil {
		helper(node.Left, path+"->"+strconv.Itoa(node.Left.Val.(int)), s)
	}
	if node.Right != nil {
		helper(node.Right, path+"->"+strconv.Itoa(node.Right.Val.(int)), s)
	}

}

func hasPathSum(root *datastructures.TreeNode, sum int) bool {
	res := binaryTreePaths(root)
	for _, path := range res {
		pathSum := 0
		for _, v := range path {
			pathSum += v
		}
		if pathSum == sum {
			return true
		}
	}
	return false
}

func main() {
	// tree1 := datastructures.NewBinaryTree(1)
	// tree1.Root.Left = datastructures.NewTreeNode(2)
	// tree1.Root.Right = datastructures.NewTreeNode(1)

	tree2 := datastructures.NewBinaryTree(1)
	tree2.Root.Left = datastructures.NewTreeNode(2)
	tree2.Root.Right = datastructures.NewTreeNode(2)
	tree2.Root.Left.Left = datastructures.NewTreeNode(1)
	tree2.Root.Left.Left.Left = datastructures.NewTreeNode(1)
	tree2.Root.Left.Left.Left.Left = datastructures.NewTreeNode(1)

	// fmt.Println(isSameTree(tree1.Root, tree2.Root))
	// fmt.Println(levelOrder(tree2.Root))
	// fmt.Println(maxDepth(tree2.Root))
	// fmt.Println(minDepth(tree2.Root))

	// preorder := []int{3, 9, 20, 15, 7}
	// inorder := []int{9, 3, 15, 20, 7}
	// buildTreeFromPreAndInOrder(preorder, inorder)

	fmt.Println(hasPathSum(tree2.Root, 6))
}
