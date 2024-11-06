package main

import "fmt"

func enigma(a ***int, b *int, c *******int, d ****int) {
	temp_a := ***a
	temp_b := *b
	temp_c := *******c
	temp_d := ****d

	// swap 3la 7sab task :
	// a ndiroha f'c
	***a = temp_c
	// b ndiroha f'a
	*b = temp_a
	// c ndiroha f'd
	*******c = temp_d
	// d ndiroha f'b
	****d = temp_b
}

func main() {
	// n'swapiw a b'3 dyal les layers
	x := 5
	y := &x
	z := &y
	a := &z

	// n'swapiw b b'layer wa7ed
	w := 2
	b := &w

	// n'swapiw c b'7 dyal les layers
	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	// n'swapiw d b'4 dyal les layers
	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	// print bla pointer
	fmt.Println("before using enigma:")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	enigma(a, b, c, d)

	// avec pointers
	fmt.Println("after using enigma:")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
