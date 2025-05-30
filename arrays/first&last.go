package arrays

import "fmt"

func FirstAndLast() {
	a := []int{1, 3, 5, 5, 5, 5, 67, 123, 125}
	target := 5
	l := bSearch(a, target, true)
	r := bSearch(a, target, false)
	fmt.Printf("%-30s %d\n", "Found at", []int{l, r})
}

func bSearch(a []int, target int, findStart bool) int {
	index := -1
	low, high := 0, len(a)-1
	for low <= high {
		m := high - (high-low)/2
		if a[m] == target {
			index = m
			if findStart {
				high = m - 1
			} else {
				low = m + 1
			}
		} else if a[m] < target {
			low = m + 1
		} else {
			high = m - 1
		}
	}
	return index
}
