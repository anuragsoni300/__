package common

import "fmt"

func GCD() {
	fmt.Printf("%-60s %d\n", "Searched in rotated array found at", _gcd(20, 24))
}

func _gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return _gcd(b, a%b)
}
