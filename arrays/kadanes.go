package arrays

import (
	"fmt"
	"math"
)

func Kadanes() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	currSum := 0
	ans := math.MinInt
	for _, v := range a {
		currSum += v
		ans = max(currSum, ans)
		if currSum < 0 {
			currSum = 0
		}
	}
	fmt.Printf("%-60s %v\n", "Two sum index", ans)
}
