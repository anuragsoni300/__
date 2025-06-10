package arrays

import "fmt"

func ZigZag() {
	a := []int{4, 3, 7, 8, 6, 2, 1}
	flag := true
	for i := range len(a) - 1 {
		if flag {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		} else {
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		flag = !flag
	}
	fmt.Printf("%-60s %v\n", "ZigZag", a)
}
