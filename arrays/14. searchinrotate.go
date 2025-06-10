package arrays

import "fmt"

func SearchInRotateArr() {
	a := []int{9, 10, 1, 2, 3, 4, 5, 6, 7, 8}
	target, ans := 1, -1
	l, r := 0, len(a)-1
	for l < r {
		m := l + (r-l)/2
		if a[m] == target {
			ans = m
			break
		}
		if a[l] < a[m] {
			if target >= a[l] && target < a[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > a[m] && target <= a[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	fmt.Printf("%-60s %d\n", "Searched in rotated array found at", ans)
}
