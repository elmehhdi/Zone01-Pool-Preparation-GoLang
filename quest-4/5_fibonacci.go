package main

import "fmt"

func fibonacci(index int) int {
	if index < 0 {
		return -1 // nreturniw -1 f negative numbers
	}
	if index == 0 {
		return 0 // fibo lowel
	}
	if index == 1 {
		return 1 // fibo zawj
	}

	// la recurssion
	return fibonacci(index-1) + fibonacci(index-2)
}

func main() {
	x := fibonacci(10)
	fmt.Println(x)
}
