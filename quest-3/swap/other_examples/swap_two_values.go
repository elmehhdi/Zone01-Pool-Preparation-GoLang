package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 5
	y := 10
	swap(&x, &y)                  // Pass the memory addresses of x and y
	fmt.Println("x:", x, "y:", y) // x: 10, y: 5
}
