package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}

func getLinkNode(data int) *LinkNode {
	return &LinkNode{data, nil}
}

type SingleLL struct {
	head *LinkNode
}

func getSingleLL() *SingleLL {
	return &SingleLL{nil}
}

func (this *SingleLL) addNode(value int) {

	var node *LinkNode = getLinkNode(value)
	if this.head == nil {
		this.head = node
	} else {
		var temp *LinkNode = this.head

		for temp.next != nil {

			temp = temp.next
		}

		temp.next = node
	}
}

func (this SingleLL) show() {
	if this.head != nil {
		var temp *LinkNode = this.head
		for temp != nil {

			fmt.Print("  ", temp.data)

			temp = temp.next
		}
	} else {
		fmt.Print("Empty\n")
	}
}

func (this SingleLL) mid() {
	if this.head == nil {

		fmt.Println("Empty")
	} else {
		var temp *LinkNode = this.head
		var mid *LinkNode = this.head

		for temp != nil && temp.next != nil &&
			temp.next.next != nil {
			mid = mid.next
			temp = temp.next.next
		}

		fmt.Println("\n Middle value : ", mid.data)
	}
}
func main() {
	var sll *SingleLL = getSingleLL()
	sll.addNode(2)
	sll.addNode(5)
	sll.addNode(1)
	sll.addNode(3)
	sll.show()
	sll.mid()
}
