package main

import "fmt"

func iterativepower(num int, power int) int {
	// ila l'2os kan negative n'returniw zero
	if power < 0 {
		return 0
	}

	result := 1

	// nmultipliyiw num m3a rasha le nombre dyal power ze3ma l'2os
	for i := 0; i < power; i++ {
		result *= num
	}

	return result
}

// ntéstew la fonction
func main() {
	x := iterativepower(4, 2)
	fmt.Println(x)
}
