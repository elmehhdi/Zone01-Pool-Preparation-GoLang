package main

import "fmt"

func isalpha(s string) bool {
	// n'loopiw 3la ga3 les charactere f string
	for _, char := range s {
		// // ila kan char chi 7aja men ghir A l Z n'returniw false
		// if kharej had range returni false
		if !(char >= 'A' && char <= 'Z') {
			return false
		}
	}
	// ila kano ga3 les characteres majuscule n'returniw true
	return true
}

func main() {
	x := isalpha("ABC")
	fmt.Println(x)
	y := isalpha("ABcde")
	fmt.Println(y)
	z := isalpha("A123")
	fmt.Println(z)
}
