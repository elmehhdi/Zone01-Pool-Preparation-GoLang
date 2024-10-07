package main

import "fmt"

func tolower(s string) string {
	res := ""
	for _, letter := range s {
		// ntchékew wach char rah majuscule bin A et Z ; bin 65 et 90
		if letter >= 'A' && letter <= 'Z' {
			// ila kan majuscule nzido 3lih 32 bach n'convértéweh l khoh minuscule
			res += string(letter + 32)

		} else {
			// si nn ila kan char minuscule n'kheliweh kima rah
			res += string(letter)
		}
	}
	return res

}

func main() {
	x := tolower("HELLO")
	fmt.Println(x)
	y := tolower("Hi I'm from MoroCCo")
	fmt.Println(y)

}
