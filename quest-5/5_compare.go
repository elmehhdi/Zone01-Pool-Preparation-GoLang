package main

import "fmt"

// f had task rana ncompariw l'unicode dyal kol charactere
// men a l z lowercase ; l'unicode howa men 97 l 122
// men A l Z uppercase ; l'unicode howa men 65 l 90
// fonction fiha deux parametres de type string
func compare(a, b string) int {
	if a == b {
		// ila kano les two strings b7al b7al n'returniw 0
		return 0
	} else if a < b {
		// ila kan string lowel 9el men string zawj n'returniw -1
		return -1
	} else {
		// ila kan string lowel kbir 3la string zawj n'returniw 1
		return 1
	}
}

func main() {
	x := compare("hi", "hi")
	fmt.Println(x)
	y := compare("ab", "abc")
	fmt.Println(y)
	z := compare("z", "a")
	fmt.Println(z)
}
