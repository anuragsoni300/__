package arrays

import "fmt"

func RotateByK() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate := 1
	rotate %= len(a)
	b := append(a[rotate:], a[:rotate]...)
	c := append(a[len(a)-rotate:], a[:len(a)-rotate]...)
	fmt.Printf("%-30s %v\n", "Rotate by K left", b)
	fmt.Printf("%-30s %v\n", "Rotate by K right", c)

	reverse(a, 0, len(a)-1)
	reverse(a, 0, rotate-1)
	reverse(a, rotate, len(a)-1)
	fmt.Printf("%-30s %v\n", "Rotate by K right O(1)", a)

	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(a, 0, len(a)-1)
	reverse(a, 0, len(a)-rotate-1)
	reverse(a, len(a)-rotate, len(a)-1)
	fmt.Printf("%-30s %v\n", "Rotate by K left O(1)", a)
}

func reverse(a []int, s, e int) {
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}
}
