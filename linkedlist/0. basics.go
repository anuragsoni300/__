package linkedlist

import "fmt"

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

	myList._print()
}
