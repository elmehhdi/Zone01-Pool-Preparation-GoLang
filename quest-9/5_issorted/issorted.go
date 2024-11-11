package main

import "fmt"

func issorted(f func(a, b int) int, a []int) bool {
	// n'loopiw 3la les elements f'slice. m3a l'false lowla n'exitew la comparaison
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

func main() {
	a1 := []int{0, 55, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	// n'compariw biha ila kan retour positive rah makayench l'order f slice
	f := func(a, b int) int {
		return a - b
	}

	result1 := issorted(f, a1)
	result2 := issorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
