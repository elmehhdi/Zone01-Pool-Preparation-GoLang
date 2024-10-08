package main

import "fmt"

func isnumeric(s string) bool {
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
	x := isnumeric("12345")
	fmt.Println(x)
	y := isnumeric("hello")
	fmt.Println(y)
	z := isnumeric("12hi33")
	fmt.Println(z)
}
