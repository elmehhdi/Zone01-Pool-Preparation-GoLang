package main

import "fmt"

func firstrune(s string) rune {
	// nholdiw s de 0 f une variable
	firstcharacter := s[0]

	// nreturniw la variable as a rune machi unicode a l'aide de la fonction rune()
	return rune(firstcharacter)
}

func main() {
	x := firstrune("hi")
	fmt.Printf("%c\n", x)

}
