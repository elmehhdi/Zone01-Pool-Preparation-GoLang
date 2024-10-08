package main

import "fmt"

func isalpha(s string) bool {
	// n'loopiw 3la ga3 les charactere f string
	for _, char := range s {
		// ila makanch char lettre men a l z wla lettre men A l Z wla ra9m men 0 l 9 nreturniw false
		if !(char >= 'A' && char <= 'Z') && !(char >= 'a' && char <= 'z') && !(char >= '0' && char <= '9') {
			return false
		}
	}
	// ila kano ga3 les charactere alphanumeric nreturniw true
	return true
}

func main() {
	x := isalpha("hi123")
	fmt.Println(x)
	y := isalpha("hi 123")
	fmt.Println(y)
	z := isalpha("hi!!")
	fmt.Println(z)
}
