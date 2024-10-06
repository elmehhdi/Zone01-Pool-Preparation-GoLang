package main

import "fmt"

func index(s string, char_to_find string) int {
	// ila kan string khawi n'returniw 0
	if char_to_find == "" {
		return 0
	}
	// n'loopiw 3la chaque letter f string
	for i := 0; i <= len(s)-len(char_to_find); i++ {
		// n'compariw wach slice equale to char li baghyino
		if s[i:i+len(char_to_find)] == char_to_find {
			return i // n'returniw la position fin l9ina char
		}
	}

	// ila mal9inach char n'returniw -1
	return -1
}

func main() {
	fmt.Println(index("hello", "l"))
	fmt.Println(index("sallam", "m"))
	fmt.Println(index("hello_world!", "w"))
	fmt.Println(index("hi", "o"))

}
