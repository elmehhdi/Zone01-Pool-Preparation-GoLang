package main

import "fmt"

func concatparams(arguments []string) string {
	result := ""

	// n'loopiw 3la slice
	for i, word := range arguments {
		result += word

		// ila makanch l'element tali f slice nzido 3lih newline
		if i < len(arguments)-1 {
			result += "\n"
		}
	}
	return result
}

func main() {
	test1 := []string{"hello", "how", "are", "you"}
	fmt.Println(concatparams(test1))
	test2 := []string{"one", "two", "three", "four", "five"}
	fmt.Println(concatparams(test2))
}
