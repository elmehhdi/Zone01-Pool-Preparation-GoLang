package main

import "fmt"

func main() {
	x := 10 // Declare an integer variable
	p := &x // Create a pointer to x

	fmt.Println("Value of x:", x)              // 10
	fmt.Println("Pointer p holds address:", p) // Memory address of x
	fmt.Println("Dereferencing p:", *p)        // Value of x, which is 10
}
