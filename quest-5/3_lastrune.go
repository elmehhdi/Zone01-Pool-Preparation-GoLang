package main

import "fmt"

func lastrune(s string) rune {
	// ila kan string khawi
	if s == "" {
		return 0
	}

	// ne3erfo toul dyal string a l'aide de la fonction len()
	length := len(s)

	// nholdiw l charactere tali f la variable (ndiro -1 la7e9ach 0 indexing)
	lastCharacter := s[length-1]

	// nconvertewha l rune w nreturniwha (la valeur li f la variable)
	return rune(lastCharacter)
}

func main() {
	fmt.Println(string(lastrune("fish")))
	fmt.Println(string(lastrune("hiii")))
	fmt.Println(string(lastrune("rain")))
	fmt.Println(string(lastrune("bird")))
	fmt.Println(string(lastrune("Solo")))
	fmt.Println(string(lastrune("menu")))
	fmt.Println(string(lastrune("music")))
	fmt.Println(string(lastrune("coatch")))
	fmt.Println(string(lastrune("pizza")))
}
