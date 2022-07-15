package main

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

func getStackNode(data int, top *StackNode) *StackNode {
	return &StackNode{
		data,
		top,
	}
}

type MyStack struct {
	top  *StackNode
	size int
}

func getMyStack() *MyStack {
	return &MyStack{
		nil,
		0,
	}
}

func (this MyStack) isSize() int {
	return this.size
}

func (this MyStack) isEmpty() bool {
	if this.isSize() > 0 {
		return false
	} else {
		return true
	}
}

func (this *MyStack) push(data int) {
	this.top = getStackNode(data, this.top)
	this.size++
}

func (this *MyStack) pop() int {
	var temp int = -1
	if !this.isEmpty() {
		temp = this.top.data
		this.top = this.top.next
		this.size--
	}
	return temp
}

func (this MyStack) peek() int {
	if !this.isEmpty() {
		return this.top.data
	} else {
		return -1
	}
}

func nextGreater(arr []int, size int) {
	var s *MyStack = getMyStack()
	for i := 0; i < size; i++ {
		if s.isEmpty() || s.peek() >= arr[i] {
			s.push(arr[i])
		} else {
			for !s.isEmpty() && s.peek() < arr[i] {
				fmt.Print(" ", s.pop())
			}
			s.push(arr[i])
		}
	}

	for !s.isEmpty() {
		fmt.Print(" ", s.pop())
	}
}

func main() {
	var arr = []int{3, 10, 4, 2, 1, 2, 6, 1, 7, 2, 9}
	var size int = len(arr)
	nextGreater(arr, size)
}
