package main

import (
	"fmt"
	"os"
)

func main() {
	// n'loopiw 3la les arguments w nbdaw men 1 la7e9ach 0 howa smiya nta3 lprogramme
	for i := 1; i < len(os.Args); i++ {
		// nprintiw les arguments
		fmt.Println(os.Args[i])
	}
}
