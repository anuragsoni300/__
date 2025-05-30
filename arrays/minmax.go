package arrays

import (
	"fmt"
	"math"
)

func MinMaxSubArray() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}

	k, sum, max := 4, 0, math.MinInt

	for i, v := range a {
		if i >= k {
			break
		}
		sum += v
		if sum > max {
			max = sum
		}
	}

	for i := k; i < len(a); i++ {
		sum -= a[i-k]
		sum += a[i]
		if sum > max {
			max = sum
		}
	}

	fmt.Printf("%-60s %v\n", "Max", max)
}
