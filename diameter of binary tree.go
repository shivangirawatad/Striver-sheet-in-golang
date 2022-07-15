package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func getTreeNode(data int) *TreeNode {
	return &TreeNode{data, nil, nil}
}

type BinaryTree struct {
	root *TreeNode
}

func getBinaryTree() *BinaryTree {

	return &BinaryTree{nil}
}

func (this BinaryTree) diameteroftree(root *TreeNode) int {
	max := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.left), dfs(node.right)
		max = Max(max, left+right)
		return 1 + Max(left, right)
	}
	dfs(root)
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	var tree *BinaryTree = getBinaryTree()

	tree.root = getTreeNode(1)
	tree.root.left = getTreeNode(2)
	tree.root.right = getTreeNode(3)
	tree.root.left.left = getTreeNode(4)
	tree.root.left.right = getTreeNode(4)
	fmt.Println(tree.diameteroftree(tree.root))
}
