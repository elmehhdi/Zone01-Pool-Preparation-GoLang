package main

import "fmt"

func max(a []int) int {
	// ila kan slice khawi n'returniw 0 3la 7sab task
	if len(a) == 0 {
		return 0
	}

	max_value := a[0]

	// n'loopiw 3la slice element b element w n'compariw
	for _, num := range a {
		if num > max_value {
			max_value = num
		}
	}

	return max_value
}

func main() {
	a := []int{10, 123, 1, 66, 23, 9}
	maximum := max(a)
	fmt.Println(maximum)
}
