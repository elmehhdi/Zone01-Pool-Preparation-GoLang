package main

import "fmt"
// 3endna deux parametres d'entree w un parametre de sortie
func division(a int, b int) int {
	c := a / b
	return c
}

func main() {
	x := 10
	y := 5
	div := division(x, y)
	fmt.Println(div)

}
