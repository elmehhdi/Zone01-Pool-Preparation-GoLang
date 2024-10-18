package main

import (
	"fmt"
	"os"
)

func even(nbr int) bool {
	return nbr%2 == 0
}

func is_even(nbr int) bool {
	if even(nbr) {
		return true
	} else {
		return false
	}
}

var even_message = "it is an even num of arguments"
var odd_message = "it is an odd num of argument"

func main() {
	// number nta3 les arguments ne9sso meno smiya dyal l'programme
	len_of_args := len(os.Args) - 1

	if is_even(len_of_args) {
		fmt.Println(even_message)
	} else {
		fmt.Println(odd_message)
	}
}
