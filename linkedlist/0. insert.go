package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type Linklist struct {
	Head *Node
}

func (list *Linklist) Insert(val int) {
	newNode := &Node{Val: val}
	if list.Head == nil {
		list.Head = newNode
	} else {
		curr := list.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
}

func (list *Linklist) InsertAfter(v1, v2 int) {
	newNode := &Node{Val: v2}

	curr := list.Head
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

func (list *Linklist) Length() {
	count := 0
	curr := list.Head
	for curr != nil {
		count++
		curr = curr.Next
	}
	fmt.Printf("%-60s %v\n", "Length of thr linked list", count)
}

func (list *Linklist) Print() {
	curr := list.Head
	for curr != nil {
		fmt.Printf("%v->", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}
