package main

import (
	"fmt"
	"os"
)

func main() {
	// ncreyew variiable njem3o fiha total dyal les arguments
	total := len(os.Args)

	// nloopiw 3la les arguments in reverse nbdaw men argument tali
	for i := total - 1; i > 0; i-- {
		fmt.Println(os.Args[i])
	}
}
