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
	root  *TreeNode
	level int
}

func getBinaryTree() *BinaryTree {
	var me *BinaryTree = &BinaryTree{}
	me.root = nil
	me.level = -1
	return me
}

func (this *BinaryTree) leftView(node *TreeNode, depth int) {
	if node != nil {
		if this.level < depth {

			fmt.Print("  ", node.data)
			this.level = depth
		}

		this.leftView(node.left, depth+1)

		this.leftView(node.right, depth-1)
	}
}

func (this *BinaryTree) rightView(node *TreeNode, depth int) {
	if node != nil {
		if this.level < depth {

			fmt.Print("  ", node.data)
			this.level = depth
		}

		this.rightView(node.right, depth+1)

		this.rightView(node.left, depth-1)
	}
}

func (this *BinaryTree) top() {
	this.level = -1
	this.leftView(this.root, 0)
	this.level = 0
	this.rightView(this.root, 0)
}

func main() {

	var tree *BinaryTree = getBinaryTree()

	tree.root = getTreeNode(1)
	tree.root.left = getTreeNode(2)
	tree.root.right = getTreeNode(4)
	tree.root.right.right = getTreeNode(5)
	tree.root.right.left = getTreeNode(6)
	tree.root.left.left = getTreeNode(7)

	tree.top()
}
