package main

import "fmt"

func printwordstables(a []string) {
	// n'loopiw 3la slice
	for _, word := range a {
		// n'printiw kol word f ster jdid 3la 7ssab task a l'aide de la fonction println()
		fmt.Println(word)
	}
}

func main() {
	test1 := []string{"hi", "ok", "one", "two"}
	printwordstables(test1)
	test2 := []string{"one", "two", "three", "four"}
	printwordstables(test2)
}
