package arrays

import "fmt"

func Duplicate() {
	a := []int{1, 5, 2, 3, 4, 6, 5}
	f := 0
	for _, v := range a {
		f = a[a[f]]

		if v == f {
			break
		}
	}
	fmt.Printf("%-60s %v\n", "Duplicate", f)
}
