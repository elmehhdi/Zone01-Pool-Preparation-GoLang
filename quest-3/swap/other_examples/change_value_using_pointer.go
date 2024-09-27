package main

import "fmt"

func main() {
	x := 20
	p := &x // Create a pointer to x

	*p = 50                           // Change the value at the memory address pointed to by p
	fmt.Println("New value of x:", x) // 50
}
