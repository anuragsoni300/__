package arrays

import "fmt"

func RotationCout() {
	a := []int{9, 10, 11, 2, 3, 4, 5, 6, 7, 8}
	l, r := 0, len(a)-1
	for l < r {
		m := l + (r-l)/2
		if a[m] > a[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	fmt.Printf("%-30s %d\n", "Rotation count", a[l])
}
