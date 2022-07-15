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
	element := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = element
	} else {
		element.Next = s.Head
		s.Head = element
	}
	s.Len++
}

func add(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev *ListNode
	var result *ListNode
	carry := 0

	for l1 != nil && l2 != nil {

		sum := l1.Val + l2.Val + carry

		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}

		tempNode := &ListNode{Val: sum}

		if prev == nil {
			result = tempNode
			prev = tempNode
		} else {
			prev.Next = tempNode
			prev = prev.Next
		}

		l1 = l1.Next

		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carry
		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}
		tempNode := &ListNode{Val: sum}
		prev.Next = tempNode
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carry
		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}
		tempNode := &ListNode{Val: sum}
		prev.Next = tempNode
		l2 = l2.Next
	}

	if carry > 0 {
		tempNode := &ListNode{Val: carry}
		prev.Next = tempNode
	}
	return result
}

func main() {
	first := initList()
	first.AddFront(1)
	first.AddFront(2)
	first.AddFront(3)

	second := initList()
	second.AddFront(1)
	second.AddFront(1)
	second.AddFront(5)

	result := add(first.Head, second.Head)
	result.Traverse()
}
