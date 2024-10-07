package main

import "fmt"

func toupper(s string) string {
	// string khawi njem3o fih resultat
	res := ""

	// n'loopiw 3la chaque char f string
	for _, letter := range s {
		// ntchékew wach char rah lowercase bin a et z ; bin 97 et 122
		if letter >= 'a' && letter <= 'z' {
			// ila kan lowercase ne9so meno 32 , 3lach ? la7e9ach bin a et A 32 characteres
			res += string(letter - 32)
		} else {
			// ila makanch dak char lowercase wla chi ponctuation n'kheliweh kima rah w nzidoh f res
			res += string(letter)
		}
	}
	return res
}
func main() {
	x := toupper("hiiii?!")
	fmt.Println(x)
}
