package main

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data int) {
	n := &Node{
		Next: nil,
		Data: data,
	}

	if l.Head == nil {
		l.Head = n
		return
	}

	curr := l.Head
	//1->2->3->nil
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
}

func (l *LinkedList) Print() {
	curr := l.Head
	//1->2->3->4
	for curr != nil {
		fmt.Printf("-> %d", curr.Data)
		curr = curr.Next
	}
	println()
}

func (l *LinkedList) Delete(data int) {
	prev := l.Head
	curr := l.Head

	for curr != nil {
		if curr.Data == data {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}
}

func main() {
	l := NewLinkedList()
	l.Insert(1)
	l.Insert(3)
	l.Insert(5)
	l.Insert(7)
	l.Print()
	l.Delete(3)
	l.Print()
}
