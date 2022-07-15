package main

import "fmt"

func initList() *SingleList {
	return &SingleList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func (s *SingleList) AddFront(num int) {
	ele := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		ele.Next = s.Head
		s.Head = ele
	}
	s.Len++
}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func main() {
	first := initList()
	first.AddFront(1)
	first.AddFront(2)
	first.AddFront(3)
	first.AddFront(4)
	first.AddFront(5)

	result := hasCycle(first.Head)
	if result {
		fmt.Println("Cycle present")
		return
	}
	//else if result == false{
	fmt.Println("Cycle not present")
	//}
}
