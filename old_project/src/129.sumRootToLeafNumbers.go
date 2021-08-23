package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) [][]int {
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

func helper(node *TreeNode, path string, s *[]string) {
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

func sumNumbers(root *TreeNode) int {
	res := binaryTreePaths(root)
	sum := 0
	for _, path := range res {
		tmpSum := 0
		for _, v := range path {
			tmpSum = 10*tmpSum + v
		}
		sum += tmpSum
	}
	return sum
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := root.Val.(int)

	var dfs func(*TreeNode) int
	// 返回，以 root 为起点，所有可能路径的 sum 值中的最大值。
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		sum := left + root.Val.(int) + right
		if maxSum < sum {
			maxSum = sum
		}

		return max(left, right) + root.Val.(int)
	}

	dfs(root)

	return maxSum
}

func testSumNumbers129() {
	tree := NewBinaryTree(-10)
	tree.Root.Left = NewTreeNode(9)
	tree.Root.Right = NewTreeNode(20)

	tree.Root.Right.Left = NewTreeNode(15)
	tree.Root.Right.Right = NewTreeNode(7)

	// tree.Root.Left.Left = NewTreeNode(3)
	// tree.Root.Left.Right = NewTreeNode(3)

	// tree.Root.Left.Left.Left = NewTreeNode(4)

	//fmt.Println(sumNumbers(tree.Root))

	fmt.Println(maxPathSum(tree.Root))
}
