package arrays

import (
	"fmt"
)

func MissingNo() {
	a := []int{8, 2, 4, 5, 3, 7, 1}

	sum, n := 0, len(a)+1
	xor1, xor2 := 0, 0
	_sum := (n) * (n + 1) / 2
	for _, v := range a {
		sum += v
	}

	for _, v := range a {
		xor1 ^= v
	}

	for i := range len(a) + 1 {
		xor2 ^= (i + 1)
	}
	fmt.Printf("%-60s %v %d\n", "Missing no.", _sum-sum, xor1^xor2)
}
