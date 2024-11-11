package main

import "fmt"

// deux parametres : fonction t'returni une boolean et un slice dyal les string
func any(f func(string) bool, a []string) bool {
	// n'loopiw 3la slice
	for _, str := range a {
		// ila t'evaluya b true n'returniw true
		if f(str) {
			return true
		}
	}
	return false
}

// fonction katchouf wach string fih chi numero si oui t'returni true
func isnumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func main() {
	str1 := []string{"hi", "hi", "6"}
	str2 := []string{"hi", "hello", "ok"}

	res1 := any(isnumeric, str1)
	res2 := any(isnumeric, str2)

	fmt.Println(res1)
	fmt.Println(res2)

}
