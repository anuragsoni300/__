package arrays

import "fmt"

func CountSubArrXork() {
	arr := []int{4, 2, 2, 6, 4}
	prefXOR := 0
	k := 6
	ans := 0
	mp := make(map[int]int, len(arr))
	for _, v := range arr {
		prefXOR ^= v
		ans += mp[prefXOR^k]
		mp[prefXOR]++
		if prefXOR == k {
			ans++
		}
	}
	fmt.Printf("%-60s %v\n", "Count of XOR subarr", ans)
}
