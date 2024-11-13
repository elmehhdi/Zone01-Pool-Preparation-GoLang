package main

import "fmt"

func advancedsortwordarr(a []string, f func(a, b string) int) {
	// l'bubble sort f' l'implementation
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// n'compariw 'a' m3a 'b' f ascci value nta3hom
func compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "e", "3", "C", "7", "5", "M", "c"}
	advancedsortwordarr(result, compare)
	fmt.Println(result)
}
