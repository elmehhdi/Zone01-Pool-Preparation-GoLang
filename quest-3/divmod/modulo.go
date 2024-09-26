package main

import "fmt"

func modulo(a int, b int) int {
	c := a % b
	return c
}

func main() {
	x := 15
	y := 2
	mod := modulo(x, y)
	fmt.Println(mod)

}
