package arrays

import (
	"fmt"
	"slices"
	"sort"
)

func EvenOddSort() {
	a := []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}

	var odd, even []int

	for i, v := range a {
		if i%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	slices.Sort(even)
	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})

	for i, v := range even {
		a[2*i] = v
	}
	for i, v := range odd {
		a[2*i+1] = v
	}
	fmt.Printf("%-60s %v\n", "Even odd sort", a)
}
