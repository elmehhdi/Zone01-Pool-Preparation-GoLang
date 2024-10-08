package main

import "fmt"

func isalpha(s string) bool {
	// n'loopiw 3la ga3 les characteres f string
	for _, char := range s {
		// ila kan char chi 7aja men ghir 0 l 9 n'returniw false
		if !(char >= '0' && char <= '9') {
			return false
		}
	}
	// ila kan string fih ghi les nums n'returniw true
	return true
}

func main() {
	x := isalpha("12345")
	fmt.Println(x)
	y := isalpha("hello")
	fmt.Println(y)
	z := isalpha("12hi33")
	fmt.Println(z)
}
