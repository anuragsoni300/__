package arrays

import "fmt"

func ZeroSumSubarray() {
	arr := []int{1, 4, -2, -2, 5, -4, 3}
	mp := make(map[int]int, len(arr))
	sum := 0

	for i, v := range arr {
		sum += v
		if _, ok := mp[sum]; ok {
			fmt.Printf("%-60s %v\n", "0 Sum Subarray Exist", ok)
			return
		}
		mp[v] = i
	}
}
