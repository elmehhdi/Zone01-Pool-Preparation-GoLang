/*
The Collatz conjecture : is one of the most famous unsolved problems in mathematics.
The conjecture asks whether repeating two simple arithmetic operations will eventually
transform every positive integer into 1. It concerns sequences of integers in which each
term is obtained from the previous term as follows: if a term is even, the next term is one half of it.
If a term is odd, the next term is 3 times the previous term plus 1.
*/

package main

import "fmt"

func collatzcountdown(start int) int {
	if start <= 0 {
		return -1
	}

	steps := 0

	// n'loopiw 7ta start tweli = 1
	for start != 1 {
		if start%2 == 0 {
			// ila start kant paire
			start /= 2
		} else {
			// ila start kant impaire
			start = start*3 + 1
		}
		steps++
	}

	return steps
}

func main() {
	num_of_steps := collatzcountdown(23)
	fmt.Println(num_of_steps)
}
