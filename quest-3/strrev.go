package main

import "fmt"

func strrev(s string) string {
	// initialize wa7ed string khawi fih ghadi t'reversé
	str := ""
	for _, rune := range s {
		// str = rune + dakchi li kan f str
		// b had l'operation ghadi yeb9a ytzad chaque rune f lbedya dyal string
		// (prepanding not appending)
		str = string(rune) + str
	}
	return str

}

func main() {
	x := "hello"
	// before reversing
	fmt.Println(x)
	y := strrev(x)
	// after the reverse
	fmt.Println(y)
}
