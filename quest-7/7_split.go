package main

import (
	"fmt"
	"strings"
)

// fonction fiha deux parametres de type string w katreturni un slice de type string
func split(str string, separateur string) []string {
	// a l'aide de la fonction split
	// syntaxe dyal split() ; syntax : Split(string li ghadi t'splitih , un separateur de type string 7ta howa li ghadi te9sem bih string)
	return strings.Split(str, separateur)
}

func main() {
	my_str := "myoknameokisokmehdi"
	separ := "ok" // ay string nsepariw bih ghadi ywli espace
	result := split(my_str, separ)

	fmt.Println(result)
}
