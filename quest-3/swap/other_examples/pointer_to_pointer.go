package main

import "fmt"

func main() {
	x := 100
	p := &x  // Pointer to x
	pp := &p // Pointer to pointer

	fmt.Println("Value of x:", x)          // 100
	fmt.Println("Dereferencing p:", *p)    // 100
	fmt.Println("Dereferencing pp:", **pp) // 100 (Double dereference)
}
