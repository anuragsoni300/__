package arrays

import "fmt"

func ContainerWithMostWater() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	l, r := 0, len(a)-1
	total := 0
	for l < r {
		total = max(total, min(a[l], a[r])*(r-l))
		if a[l] < a[r] {
			l++
		} else {
			r--
		}
	}
	fmt.Printf("%-60s %v\n", "Two sum index", total)
}
