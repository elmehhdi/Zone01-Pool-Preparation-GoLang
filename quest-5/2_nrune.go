package main

import "fmt"

func nrune(s string, n int) rune {
	// ila kan n 9el men zero nreturniw 0 logiquement string yebda f 1 (yebda f 0 language are 0 indexed)
	if n <= 0 {
		return 0
	}
	// ila kan n kter men string nreturniw 0
	if n > len(s) {
		return 0
	}

	// nreturniw n-1 3la 7ssab task
	return rune(s[n-1])
}

func main() {
	x := nrune("seven_height_nine_ten", 4)
	fmt.Printf("%c\n", x)
	y := nrune("seven_height_nine_ten", 6)
	fmt.Printf("%c\n", y)
	z := nrune("seven_height_nine_ten", 12)
	fmt.Printf("%c\n", z)
}
