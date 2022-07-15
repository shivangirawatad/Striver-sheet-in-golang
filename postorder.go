package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) postorder() {
	if root != nil {
		root.left.postorder()
		root.right.postorder()
		fmt.Printf("%d ", root.data)
	}
	return
}
func main() {
	tree := Node{2, &Node{1, &Node{3, nil, nil}, &Node{5, nil, nil}}, &Node{4, &Node{6, nil, nil}, &Node{7, nil, nil}}}
	tree.postorder()
}
