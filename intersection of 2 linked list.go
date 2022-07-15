package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}

func getLinkNode(data int) *LinkNode {

	return &LinkNode{
		data,
		nil,
	}
}

type SingleLL struct {
	head *LinkNode
	tail *LinkNode
}

func getSingleLL() *SingleLL {

	return &SingleLL{
		nil,
		nil,
	}
}
func (this *SingleLL) insert(data int) {
	var node *LinkNode = getLinkNode(data)
	if this.head == nil {

		this.head = node
	} else {

		this.tail.next = node
	}

	this.tail = node
}

func (this SingleLL) display() {
	if this.head == nil {
		return
	}
	var temp *LinkNode = this.head

	for temp != nil {
		fmt.Print(" ", temp.data, " →")

		temp = temp.next
	}
	fmt.Println(" NULL")
}

func (this SingleLL) length() int {
	var temp *LinkNode = this.head
	var counter int = 0

	for temp != nil {
		counter++

		temp = temp.next
	}
	return counter
}

func intersectionNode(
	list1 *LinkNode,
	list2 *LinkNode, n int) *LinkNode {

	var t1 *LinkNode = list1
	var t2 *LinkNode = list2

	var counter int = 0

	for t1 != nil && counter < n {

		t1 = t1.next
		counter++
	}
	if counter != n || t1 == nil {
		return nil
	} else {
		for t1 != nil && t2 != nil {
			if t1 == t2 {

				return t1
			}

			t1 = t1.next
			t2 = t2.next
		}
		return nil
	}
}

func (this SingleLL) intersectionPoint(other *SingleLL) {

	var result *LinkNode = nil

	var l1 int = this.length()
	var l2 int = other.length()

	fmt.Print("\n Linked Linked A : ")
	this.display()
	fmt.Print(" Linked Linked B : ")
	other.display()
	if l1 > 0 && l2 > 0 {
		if l1 > l2 {

			result = intersectionNode(this.head, other.head, l1-l2)
		} else {
			result = intersectionNode(other.head, this.head, l2-l1)
		}
	}
	if result == nil {
		fmt.Println(" Intersection Point are not exist")
	} else {
		fmt.Println(" Intersection Point : ", result.data)
	}
}
func main() {
	var list1 *SingleLL = getSingleLL()
	var list2 *SingleLL = getSingleLL()

	// 1 → 5 → 3 → 2 → 7 → 11 → 4 → NULL
	list1.insert(1)
	list1.insert(5)
	list1.insert(3)
	list1.insert(2)
	list1.insert(7)
	list1.insert(11)
	list1.insert(4)

	list2.insert(4)
	list2.insert(6)
	list2.insert(5)
	list2.insert(8)

	list2.tail.next = list1.head.next.next
	list1.intersectionPoint(list2)
	list2.tail.next = list1.head.next.next.next
	list1.intersectionPoint(list2)
	list2.tail.next = list1.head
	list1.intersectionPoint(list2)
	list2.tail.next = nil
	list1.tail.next = list2.head.next
	list1.intersectionPoint(list2)
	list1.tail.next = nil
}
