package main

import "fmt"

// 3endna deux parametres dentree w deux parametre de sortie de type int
func divmod(a int, b int) (int, int) {
	div := a / b
	mod := a % b
	return div, mod
}

func main() {
	x := 10
	y := 2
	div, mod := divmod(x, y)
	fmt.Println(div, mod)
}
