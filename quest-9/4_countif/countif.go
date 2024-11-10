package main

import "fmt"

func countif(f func(string) bool, tab []string) int {
	count := 0
	// n'loopiw 3la les string f tab
	for _, str := range tab {
		// ila 'f' t'evaluyat b'true nzido count b'1
		if f(str) {
			count++
		}
	}
	return count
}

func isnumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func main() {
	tab1 := []string{"hi", "how", "are", "you"}
	tab2 := []string{"hi", "5", "ok", "4", "hello", "24"}
	answer1 := countif(isnumeric, tab1)
	answer2 := countif(isnumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
