package arrays

import "fmt"

func TrappingRainWater() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	l_max, r_max, total := 0, 0, 0
	l, r := 0, len(a)-1

	for l < r {
		if a[l] <= a[r] {
			if a[l] < l_max {
				total += l_max - a[l]
			} else {
				l_max = a[l]
			}
			l++
		} else {
			if a[r] < r_max {
				total += r_max - a[r]
			} else {
				r_max = a[r]
			}
			r--
		}
	}
	fmt.Printf("%-60s %v\n", "Trapped rain water", total)
}
