package main

import "fmt"

func Map(f func(int) bool, a []int) []bool {
	// n'creyéw slice fin n'storiw result a l'aide de la fonction make
	result := make([]bool, len(a))
	// n'loopiw 3la ga3 les elements li f' a
	for i, v := range a {
		// et n'aplikiw la fonction 'f' 3la kol element
		result[i] = f(v)
	}
	return result
}

// la fonction is prime katreturniw true ila kan num prime w katreturni false ila kan num not prime
func isprime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(isprime, a)
	fmt.Println(result)
}
