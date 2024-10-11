package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	// os.Args[0] howa smiya dyal l'porogramme dakchi 3lach kanbdaw men 1
	myarguments := os.Args[1:]

	// la fonction Strings() katsorté les strings en se basant 3la ascci code nta3hom
	sort.Strings(myarguments)

	// n'loopiw 3la les arguments w nteb3o kol wa7ed
	for _, arg := range myarguments {
		fmt.Println(arg)
	}
}
