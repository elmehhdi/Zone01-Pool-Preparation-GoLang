package main

import "fmt"

func basicjoin(elems []string) string {
	// string khawi fin njem3o resultat dyal lconcatenation
	result := ""

	// n'loopiw 3la ga3 les strings
	for _, elem := range elems {
		// n'concaténéw kol string loopina 3lih f la variable result
		result += elem
	}

	return result
}

func main() {
	x := []string{"helo!", " how", " are", " you"}
	// test n'printiw fih x bla mankhedmo b la concatenation 
	fmt.Println(x)
	y := basicjoin([]string{"hello!", " how", " are", " you"})
	// avec la concatenation dyal une liste (on appelle list un slice f go)
	fmt.Println(y)
}
