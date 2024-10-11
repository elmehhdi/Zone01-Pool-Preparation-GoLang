package main

import "fmt"

// fonction fiha deux parametres slice w string fih separateur
func join(elems []string, sep string) string {

	// n'creyew variable fih l'element lowel dyal slice (liste)
	result := elems[0]

	// n'loopiw 3la les elements lokhrin f slice ; nebdaw men l'element li l'index nta3o 1 w manfotoch toul dyal slice a l'aide de la fonction len()
	for i := 1; i < len(elems); i++ {
		// nzido separateur f kol loop
		result += sep + elems[i]
	}

	// nreturniw result nta3 lconcatenation avec separateur
	return result
}

func main() {
	x := []string{"hi", "hello", "ok", "okk"}
	fmt.Println(join(x, ":"))
}
