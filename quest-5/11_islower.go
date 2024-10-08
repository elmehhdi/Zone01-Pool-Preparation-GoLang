package main

import "fmt"

func islower(s string) bool {
	// n'loopiw 3la ga3 les characteres f string
	for _, char := range s {
		// ila kan char chi 7aja men ghir a l z n'returniw false
		// if kharej had range (minuscule) returni false
		if !(char >= 'a' && char <= 'z') {
			return false
		}
	}
	// ila kano ga3 les charactere minuscule nreturniw true
	return true
}

func main() {
	x := islower("abcd")
	fmt.Println(x)
	y := islower("abCd")
	fmt.Println(y)
	z := islower("hi!!")
	fmt.Println(z)
}
