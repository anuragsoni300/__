package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type Linklist struct {
	head *Node
}

func (list *Linklist) _insert(val int) {
	newNode := &Node{Val: val}
	if list.head == nil {
		list.head = newNode
	} else {
		curr := list.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
}

func (list *Linklist) _insertAfter(v1, v2 int) {
	newNode := &Node{Val: v2}

	curr := list.head
	for curr != nil {
		if curr.Val == v1 {
			break
		}
		curr = curr.Next
	}
	temp := curr.Next
	curr.Next = newNode
	newNode.Next = temp
}

func (list *Linklist) _delete(val int) {
	curr := list.head
	prev := curr
	for curr != nil && curr.Val != val {
		prev = curr
		curr = curr.Next
	}
	prev.Next = curr.Next
}

func (list *Linklist) _reverse() {
	curr := list.head
	var prev *Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

func (list *Linklist) _print() {
	curr := list.head
	for curr != nil {
		fmt.Printf("%v->", curr.Val)
		curr = curr.Next
	}
}

func Linkedlist() {
	myList := Linklist{}
	myList._insert(1)
	myList._insert(2)
	myList._insert(3)
	myList._insert(4)

	myList._insertAfter(2, 5)
	myList._delete(3)
	myList._print()
	myList._reverse()
	myList._print()
}
