package main

import "fmt"

// fonction 3endha parametre wa7ed kaypointi 3la wa7ed l int w ychanger la valeur dyalo l 6
func PointOne(n *int) {
	*n = 6 // "*" l'operateur de dereferencement
}

func main() {
	num := 10
	// printina la valeure dyal num 3adi
	fmt.Println(num)
	// changena la valeur dyal num en passant par l'adressse dyalha
	PointOne(&num) // "&" pass the value of num l function
	// printina la valeur jdida nta3 num
	fmt.Println(num)
}

// Un pointeur est une variable qui stocke l'adresse mémoire
// d'une autre variable.
