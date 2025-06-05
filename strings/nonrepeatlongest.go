package strings

import "fmt"

func LongestNonRepeting() {
	a := "pwwkew"
	mp := make(map[rune]int)
	ans, lastIndex := 0, 0
	for i, rn := range a {
		if index, ok := mp[rn]; ok && index >= lastIndex {
			lastIndex = index + 1
		}
		mp[rn] = i
		ans = max(ans, i-lastIndex+1)
	}
	fmt.Printf("%-60s %d\n", "Length of longest sbstring non repeating", ans)
}
