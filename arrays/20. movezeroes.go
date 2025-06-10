package arrays

import "fmt"

func MoveZeroes() {
	a := []int{0, 10, 1, 0, 0, 0, 3, 12}

	lastNonZero := 0

	for i := range a {
		if a[i] != 0 {
			a[lastNonZero], a[i] = a[i], a[lastNonZero]
			lastNonZero++
		}
	}
	fmt.Printf("%-60s %v\n", "Move zeroes", a)
}
