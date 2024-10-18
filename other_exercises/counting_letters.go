package main

import (
	"fmt"
)

func countletters(str string) {
	for i := 0; i < len(str); i++ {
		count := 0
		for j := 0; j < len(str); j++ {
			// n'compariw wach i b7al j si oui n'incrementéw count b 1
			if str[i] == str[j] {
				count++
			}
		}
		fmt.Printf("%c : %d\n", str[i], count)
	}
}

func main() {
	countletters("hi")
}
