package main

import "fmt"

func iterativefactorial(num int) int {
	if num < 0 {
		return 0 // l'factoriel nta3 num negative = 0
	}
	if num == 0 {
		return 1 // l'factoriel nta3 0 = 1
	}

	result := 1
	for i := 1; i <= num; i++ {
		result *= i // n'multipliyiw result f kol num dyal l'itération
	}
	return result // nreturniw resultat
}

func main() {
	a := iterativefactorial(5)
	fmt.Println(a)
	b := iterativefactorial(10)
	fmt.Println(b)

}
