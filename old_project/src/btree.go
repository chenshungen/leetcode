package main

import "fmt"

// TreeNode B树节点
type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode ...
func NewTreeNode(data interface{}) *TreeNode {
	return &TreeNode{Val: data}
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", node.Val, node.Left, node.Right)
}

// BinaryTree ...
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree ...
func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewTreeNode(rootV)}
}

// InOrderTraverse 中序遍历 (left->root->right)
func (btree *BinaryTree) InOrderTraverse() {
	p := btree.Root
	s := NewArrayStack()
	for !s.IsEmpty() || p != nil {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*TreeNode)
			fmt.Printf("%+v ", tmp.Val)
			p = tmp.Right
		}
	}
	fmt.Println()
}

// InOrderTraverseRecursive 递归实现
func (btree *BinaryTree) InOrderTraverseRecursive() {
	inOrderTraverseRecursive(btree.Root)
}

func inOrderTraverseRecursive(p *TreeNode) {
	if p == nil {
		return
	}
	inOrderTraverseRecursive(p.Left)
	fmt.Printf("%v ", p.Val)
	inOrderTraverseRecursive(p.Right)
}

// PreOrderTraverse 前序遍历 root->left->right
func (btree *BinaryTree) PreOrderTraverse() {
	p := btree.Root
	s := NewArrayStack()
	for !s.IsEmpty() || p != nil {
		if p != nil {
			fmt.Printf("%v ", p.Val)
			s.Push(p)
			p = p.Left
		} else {
			p = s.Pop().(*TreeNode).Right
		}
	}
	fmt.Println()
}

// PreOrderTraverseRecursive ...
func (btree *BinaryTree) PreOrderTraverseRecursive() {
	preOrderTraverseRecursive(btree.Root)
}

func preOrderTraverseRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v", root.Val)
	preOrderTraverseRecursive(root.Left)
	preOrderTraverseRecursive(root.Right)
}

// PostOrderTraverse 后序遍历left->right->root
func (btree *BinaryTree) PostOrderTraverse() {
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(btree.Root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*TreeNode)
		s2.Push(p)
		if p.Left != nil {
			s1.Push(p.Left)
		}
		if p.Right != nil {
			s1.Push(p.Right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%v ", s2.Pop().(*TreeNode).Val)
	}
}

// PostOrderTraverseRecursive 后序遍历 递归实现
func (btree *BinaryTree) PostOrderTraverseRecursive() {
	postOrderTraverseRecursive(btree.Root)
}

func postOrderTraverseRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraverseRecursive(root.Left)
	postOrderTraverseRecursive(root.Right)
	fmt.Printf("%v", root.Val)
}
