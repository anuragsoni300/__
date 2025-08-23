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

func (list *Linklist) _length() {
	count := 0
	curr := list.head
	for curr != nil {
		count++
		curr = curr.Next
	}
	fmt.Printf("%-60s %v\n", "Length of thr linked list", count)
}

func (list *Linklist) _print() {
	curr := list.head
	for curr != nil {
		fmt.Printf("%v->", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func Linkedlist() {
	myList, list2 := Linklist{}, Linklist{}
	myList._insert(1)
	// myList._insert(2)
	myList._insert(3)
	// myList._insert(4)
	myList._insert(5)
	// myList._insert(6)
	myList._insert(7)
	// myList._insert(8)
	myList._insert(9)
	// myList._insert(10)

	list2._insert(2)
	list2._insert(4)
	list2._insert(6)
	list2._insert(8)

	list2._insert(10)
	// myList._insertAfter(2, 5)
	// myList._delete(3)
	// myList._print()
	// myList.head = myList.head._reverseRec()
	// myList._reverse()
	// myList._print()
	// myList._length()
	// myList._middle()

	// curr := myList.head
	// for curr.Next != nil {
	// 	curr = curr.Next
	// }
	// curr.Next = myList.head.Next
	// myList._detectcycle()
	// myList._print()
	// myList.head = myList.head._mergeSortedRec(list2.head)
	myList.head = myList.head._mergeSortedItr(list2.head)
	myList._print()
}
