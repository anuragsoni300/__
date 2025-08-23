package main

import (
	"github.com/anuragsoni300/__/arrays"
	"github.com/anuragsoni300/__/common"
	"github.com/anuragsoni300/__/linkedlist"
	"github.com/anuragsoni300/__/strings"
)

func main() {
	LinkedList()
}

func Common() {
	common.GCD()
}

func Strings() {
	strings.LongestNonRepeting()
}

func Array() {
	arrays.Insertion()
	arrays.Deletion()
	arrays.BSearch()
	arrays.BSort()
	arrays.SSort()
	arrays.ISort()
	arrays.MSort()
	arrays.QSort()
	arrays.RotateByK()
	arrays.MinMaxSubArray()
	arrays.LongestSubArraySumK()
	arrays.PeakElement()
	arrays.FirstAndLast()
	arrays.RotationCout()
	arrays.SearchInRotateArr()
	arrays.TwoSum()
	arrays.RemoveDuplicates()
	arrays.PartitionArr()
	arrays.Kadanes()
	arrays.Frequencies()
	arrays.MoveZeroes()
	arrays.EvenOddSort()
	arrays.ColorSort()
	arrays.MissingNo()
	arrays.Majority()
	arrays.Duplicate()
	arrays.Equilibrium()
	arrays.TrappingRainWater()
	arrays.ContainerWithMostWater()
	arrays.MatrixSpiral()
	arrays.DiagonalTraversal()
	arrays.ZigZag()
	arrays.RotateImage()
	arrays.ZeroSumSubarray()
	arrays.CountSubArrXork()
	arrays.LongestConsecutive()
}

func LinkedList() {
	List1, List2 := linkedlist.Linklist{}, linkedlist.Linklist{}
	List3, List4 := linkedlist.Linklist{}, linkedlist.Linklist{}

	List1.Insert(1)
	List1.Insert(3)
	List1.Insert(5)
	List1.Insert(7)
	List1.Insert(9)

	List2.Insert(2)
	List2.Insert(4)
	List2.Insert(6)
	List2.Insert(8)
	List2.Insert(10)

	List3.Insert(1)
	List3.Insert(3)
	List3.Insert(5)
	List3.Insert(7)
	List3.Insert(9)

	List4.Insert(2)
	List4.Insert(4)
	List4.Insert(6)
	List4.Insert(8)
	List4.Insert(10)

	List2.InsertAfter(2, 5)
	List2.Print()

	List1.Delete(3)
	List1.Print()

	List2.Reverse()
	List2.Print()

	List2.Length()
	List2.Middle()

	List3.Head = List3.Head.MergeSortedRec(List4.Head)
	List3.Print()

	curr := List1.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = List1.Head.Next
	List1.Detectcycle()
	List1.Print()
}
