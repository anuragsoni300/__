package arrays

import "fmt"

func LongestConsecutive() {
	arr := []int{36, 41, 56, 35, 44, 33, 34, 92, 43, 32, 42}
	mp := make(map[int]bool)
	ans := 0
	for _, v := range arr {
		mp[v] = true
	}

	for _, v := range arr {
		if !mp[v-1] {
			currNum := v
			currAns := 1
			for mp[currNum+1] {
				currNum++
				currAns++
			}
			ans = max(ans, currAns)
		}
	}
	fmt.Printf("%-60s %v\n", "Longest Consecutive arr length", ans)
}
