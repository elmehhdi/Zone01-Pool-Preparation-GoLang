package main

import "fmt"

func recursivefactorial(num int) int {
	// les cas dyal invalid number
	if num < 0 {
		return 0
	} else if num == 0 {
		return 1
	} else {
		return num * recursivefactorial(num-1) // recursive call ; kat3awed telgha l rasha
	}
}

func main() {
	a := 3
	result := recursivefactorial(a)
	fmt.Println(result)
}
