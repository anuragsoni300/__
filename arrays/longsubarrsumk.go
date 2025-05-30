package arrays

import (
	"fmt"
)

func LongestSubArraySumK() {
	a := []int{-5, 8, -14, 2, 4, 12}
	target := -5
	ans, prefSum := 0, 0
	mp := make(map[int]int, len(a)+1)
	for i, v := range a {
		prefSum += v
		if prefSum == target {
			ans = i + 1
		}

		if _, ok := mp[prefSum-target]; ok {
			ans = max(ans, i-mp[prefSum-target]+1)
		}
		if _, ok := mp[prefSum]; !ok {
			mp[prefSum] = i
		}
	}
	fmt.Printf("%-30s %v\n", "Length of max subarr", ans)
}
