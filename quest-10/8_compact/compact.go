package main

import "fmt"

func compact(ptr *[]string) int {
	// n'creyew slice khawi njem3o fih ga3 les elements
	var compacted []string

	// n'loopiw 3la les elements f l'originale slice
	for _, str := range *ptr {
		// ila kan l'element 3amer n'appendiweh f slice
		if str != "" {
			compacted = append(compacted, str)
		}
	}

	// n'remplaciw slice originale b'compacted a travers l'pointeur
	*ptr = compacted

	return len(compacted)
}

const N = 10

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"
	// 9bel compact
	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("the size after compacting:", compact(&a))
	// mor compact
	for _, v := range a {
		fmt.Println(v)
	}
}
