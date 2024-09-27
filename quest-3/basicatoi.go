package main

import "fmt"

func basicatoi(s string) int {
	// initialize result
	result := 0
	// n'itériw 3la chaque rune f string
	for _, rune := range s {
		num := int(rune - '0')   // la valeur dyal 'O' f ASCII hiya 48 , par exemple rune lowel f string howa 1, la valeur dyal 1 f ASCII hiya 49 , 49-48 = 1 w canconvetew had '1' a l'aide de la fonction int().
		result = result*10 + num // nbuildiw result num b num
	}
	return result
}

func main() {
	digit := basicatoi("123")

	fmt.Println(digit + 2)
}
