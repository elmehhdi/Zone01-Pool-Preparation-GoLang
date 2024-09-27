package main

import "fmt"

func strlen(s string) int {
	count := 0
	// maghadich ne7tajo l'index (position dyal charactere)
	// w maghadich ne7tajo rune (l'charactere a chaque iteration)
	// kol iteration la variable count katzid b 1 w f tali we return it
	for range s {
		count++
	}
	return count

}

func main() {
	x := strlen("hello hi")
	fmt.Println(x)

}
