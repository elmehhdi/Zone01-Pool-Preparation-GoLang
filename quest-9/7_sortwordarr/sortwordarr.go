// bubble sort algo
package main

import "fmt"

func sortwordarr(a []string) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// n'compariw ascii value dyal a[j] et l'element li moraha w n'swapiw ila t'evaluyat b true
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}

	sortwordarr(result)
	fmt.Println(result)
}
