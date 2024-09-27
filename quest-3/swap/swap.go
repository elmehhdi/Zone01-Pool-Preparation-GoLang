package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c

}

func main() {
	a := 5
	b := 10
	// before the swap
	fmt.Println(a)
	fmt.Println(b)
	swap(&a, &b)
	// after the swap
	fmt.Println(a)
	fmt.Println(b)
}
