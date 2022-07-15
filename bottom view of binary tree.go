package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func getTreeNode(data int) *TreeNode {
	return &TreeNode{data, nil, nil}
}

type Pairs struct {
	level int
	value int
}

func getPairs(level int, value int) *Pairs {

	return &Pairs{
		level,
		value,
	}
}

type BinaryTree struct {
	root *TreeNode
	min  int
}

func getBinaryTree() *BinaryTree {
	return &BinaryTree{
		nil,
		math.MaxInt64,
	}
}

func (this *BinaryTree) getBottomView(node *TreeNode,
	distance int, level int, record map[int]*Pairs) {
	if node != nil {
		if _, found := record[distance]; found {
			if record[distance].level < level {
				var p *Pairs = record[distance]

				p.value = node.data

				p.level = level
			}
		} else {

			record[distance] = getPairs(level, node.data)
		}
		if this.min > distance {
			this.min = distance
		}
		this.getBottomView(node.left, distance-1, level+1, record)
		this.getBottomView(node.right, distance+1, level+1, record)
	}
}

func (this BinaryTree) printBottomView() {

	this.min = math.MaxInt64

	var record = make(map[int]*Pairs)
	this.getBottomView(this.root, 0, 0, record)

	var i = 0

	for i < len(record) {
		fmt.Print("  ", record[this.min].value)
		this.min += 1
		i += 1
	}
}

func main() {

	var tree *BinaryTree = getBinaryTree()

	tree.root = getTreeNode(1)
	tree.root.left = getTreeNode(2)
	tree.root.right = getTreeNode(4)
	tree.root.right.right = getTreeNode(5)
	tree.root.right.left = getTreeNode(6)
	tree.root.left.left = getTreeNode(7)

	tree.printBottomView()
}
