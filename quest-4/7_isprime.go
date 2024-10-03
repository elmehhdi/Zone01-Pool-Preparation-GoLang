package main

import "fmt"

func isprime(num int) bool {
	// la condition lowla ila chi num 9el men 2 rah not prime
	if num < 2 {
		return false
	}

	// ila kan num equal l' 2 rah prime
	if num == 2 {
		return true
	}

	// ila chi num modulo 2 = 0 rah machi prime ga3 les nombres pair machi prime
	if num%2 == 0 {
		return false
	}

	// ntchekew l9awasim men 3 7ta l jid3 (square root) dyal num
	// 3lach derna i += 2 ? déja tchékéna les nombre pair f la condition li 9bel
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false // ila kan num divisible b i so its not prime
		}
	}

	// ila mal9ina 7ta chi 9asem nreturniw true
	return true
}

func main() {
	x := isprime(10)
	fmt.Println(x)
	y := isprime(16)
	fmt.Println(y)
	z := isprime(7)
	fmt.Println(z)

}
