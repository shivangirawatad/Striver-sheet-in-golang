package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func newNode(value int, next *Node) *Node {
	var n Node
	n.value = value
	n.next = next
	return &n
}
func traverse(head *Node) {
	fmt.Printf("Linked List is: ")
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}
func reverse(head *Node) {
	if head == nil {
		return
	}
	reverse(head.next)
	fmt.Printf("%d ", head.value)
}
func main() {
	head := newNode(40, newNode(30, newNode(10, newNode(40, nil))))
	traverse(head)
	fmt.Printf("Reversed linked list: ")
	reverse(head)
}
