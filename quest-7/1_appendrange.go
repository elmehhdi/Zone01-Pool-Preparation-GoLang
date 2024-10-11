package main

import "fmt"

// fonction fiha deux parametres et return slice de int
func appendrange(min int, max int) []int {
	// netchekew ila kan l'minimum kter wla ysawi l'maximum
	if min >= max {
		// si oui n'returniw nil 3la 7sab task
		return nil
	}

	// n'initializiw wa7ed slice khawi
	result := []int{}

	// nloopiw men min 7ta l max-1
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	fmt.Println(appendrange(17, 44))
	fmt.Println(appendrange(3, 31))
	fmt.Println(appendrange(5, 5))
	fmt.Println(appendrange(10, 7))
}
