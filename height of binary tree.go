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

func (this BinaryTree) height(node *TreeNode) int {
	if node != nil {
		var a int = this.height(node.left)
		var b int = this.height(node.right)

		if a > b {
			return a + 1
		} else {
			return b + 1
		}
	} else {
		return 0
	}
}
func main() {

	var tree *BinaryTree = getBinaryTree()

	tree.root = getTreeNode(1)
	tree.root.left = getTreeNode(2)
	tree.root.right = getTreeNode(3)
	tree.root.left.left = getTreeNode(4)
	tree.root.left.left = getTreeNode(5)
	fmt.Println(tree.height(tree.root))
}
