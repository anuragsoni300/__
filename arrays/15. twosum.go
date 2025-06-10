package arrays

import "fmt"

func TwoSum() {
	a := []int{7, 8, 9, 4, 5, 3, 1, 2, 6}
	mp := make(map[int]int, len(a))
	target := 10
	ans := make([]int, 2)
	for i, v := range a {
		if j, ok := mp[target-v]; ok {
			ans = []int{j, i}
			break
		} else {
			mp[v] = i
		}
	}
	fmt.Printf("%-60s %v\n", "Two sum index", ans)
}
