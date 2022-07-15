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

func (this *BinaryTree) printLeftView(node *TreeNode, depth int) {
	if node != nil {
		if this.level < depth {
			fmt.Print("  ", node.data)
			this.level = depth
		}
		this.printLeftView(node.left, depth+1)
		this.printLeftView(node.right, depth+1)
	}
}

func (this *BinaryTree) leftView() {
	this.level = -1
	this.printLeftView(this.root, 0)
}

func main() {

	var tree *BinaryTree = getBinaryTree()

	tree.root = getTreeNode(1)
	tree.root.left = getTreeNode(2)
	tree.root.right = getTreeNode(4)
	tree.root.right.right = getTreeNode(5)
	tree.root.right.left = getTreeNode(6)
	tree.root.left.left = getTreeNode(7)

	tree.leftView()
}
