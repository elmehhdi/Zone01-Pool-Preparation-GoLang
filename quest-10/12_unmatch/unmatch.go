package main

import "fmt"

func unmatch(a []int) int {

	// loop west loop n'compariw les elements li f slice ila kan un match ntel3o count b 1
	for _, num := range a {
		count := 0
		for _, elem := range a {
			if elem == num {
				count++
			}
		}
		// ila kan l'count 1 cad ma3endoch khoh n'returniw dak num
		if count == 1 {
			return num
		}
	}
	return -1
}

func main() {
	sliceofnum := []int{5, 6, 5, 5, 6, 10}
	ok := unmatch(sliceofnum)
	fmt.Println(ok)
}
