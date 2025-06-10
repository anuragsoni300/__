package arrays

import "fmt"

func Majority() {
	a := []int{1, 4, 3, 4, 4}
	curr, count := 0, 0
	for i := range a {
		if count == 0 {
			curr = a[i]
		}
		if curr == a[i] {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, v := range a {
		if v == curr {
			count++
		}
	}

	if count <= len(a)/2 {
		curr = -1
	}
	fmt.Printf("%-60s %v\n", "Majority element", curr)
}
