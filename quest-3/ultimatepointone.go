package main

import "fmt"

func ultimatepointone(num ***int) {
	***num = 1

}

func main() {
	a := 10 // variable katstoké la valeur 10
	// printina la valeur dyal a
	fmt.Println(a)
	b := &a // pointeur fih ladresse memoire dyal a
	c := &b // pointeur fih ladresse memoire dyal b

	ultimatepointone(&c) // changena la valeur dyal a
	fmt.Println(a)

}
