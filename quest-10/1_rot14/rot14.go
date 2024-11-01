package main

import "fmt"

func rot14(s string) string {
	result := ""
	// n'loopiw 3la string
	for _, char := range s {
		// ila kan l'charactere minuscule
		if char >= 'a' && char <= 'z' {
			// nzido 14 position w ila depassina z ne9so 26 bach n3awdo dora
			new_char := char + 14
			if new_char > 'z' {
				new_char -= 26
			}
			// nbniw f result charactere b charactere
			result += string(new_char)
			// la meme chose bnesba l majuscule
		} else if char >= 'A' && char <= 'Z' {
			new_char := char + 14
			if new_char > 'Z' {
				new_char -= 26
			}
			result += string(new_char)
			// ila kan chi char machi alphabet nkheliweh kima rah
		} else {
			result += string(char)
		}
	}

	return result
}

func main() {
	str := rot14("hi")
	fmt.Println(str)
}
