package main

import "fmt"

func activebits(n int) int {
	count := 0

	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}

	return count
}

func main() {
	fmt.Println(activebits(15))
	fmt.Println(activebits(3))
	fmt.Println(activebits(7))

}
