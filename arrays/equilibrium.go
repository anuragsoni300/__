package arrays

import "fmt"

func Equilibrium() {
	a := []int{1, 7, 3, 6, 5, 6}
	sufSum := 0
	for _, v := range a {
		sufSum += v
	}

	prefSum := 0
	for i, v := range a {
		sufSum -= v
		if sufSum == prefSum {
			fmt.Printf("%-60s %v\n", "Two sum index", i)
		}
		prefSum += v
	}
}
