package main

import "fmt"

func makerange(min int, max int) []int {
	// n7essbo toul nta3 slice
	toul := max - min

	// ila kan toul nta3 slice 9el wla equal to zero n'returniw nil
	if toul <= 0 {
		return nil
	}

	// n'creyew slice avec la fonction make() 3la 7sab task , append memnou3a
	// syntax nta3 make : make(type, length, capacity). b'nesba l silce kane7tajo ghi type et toul.
	myslice := make([]int, toul)

	for i := 0; i < toul; i++ {
		myslice[i] = min + i
	}
	return myslice
}

func main() {
	test1 := makerange(4, 9)
	fmt.Println(test1)
	test2 := makerange(6, 45)
	fmt.Println(test2)
	test3 := makerange(32, 4)
	fmt.Println(test3)
}
