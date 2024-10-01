package main

import "fmt"

func sqrt(num int) int {
	if num < 0 {
		return 0
	}

	for i := 0; i <= num; i++ {
		// ila kant i * i == num , l9ina square b les instructions dyal task
		if i*i == num {
			return i
		}
	}
	return 0
}

func main() {
	x := sqrt(100)
	fmt.Println(x)
}
