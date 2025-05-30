package arrays

import "fmt"

func PeakElement() {
	a := []int{11, 8, 7, 6, 5, 3, 1, 2, 1}

	l, r := 0, len(a)-1
	for l < r {
		m := l + (r-l)/2
		if a[m] > a[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	fmt.Printf("%-30s %v\n", "Peak Element", a[l])
}
