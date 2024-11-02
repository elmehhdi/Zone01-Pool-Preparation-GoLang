package main

import (
	"fmt"
	"os"
)

func main() {
	all_alert_word := []string{"01", "galaxy", "galaxy 01"}

	// nebdaw men l'index 1 . 0 howa smiya dyal l'programme
	for _, arg := range os.Args[1:] {
		// n'loopiw ila l9ina match m3a slice n'returniw alert
		for _, word := range all_alert_word {
			if arg == word {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
