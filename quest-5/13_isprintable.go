package main

import "fmt"

func isprintable(s string) bool {
	// n'loopiw 3la ga3 les characteres f string
	for _, char := range s {
		// range dyal l'printable charactere bin 32 et 126
		// l'printable charactere homa li kter men 32 (unicode dyalhom) wla 9el men 126 ; 9el men 32 smiythom control charactere w fog 126 smiythom non printable
		// n'returniw false ila oui
		if char < 32 || char > 126 {
			return false
		}
	}
	// n'returniw true ila kano les chractere printable 3la 7ssab task
	return true
}

func main() {
	x := isprintable("hi")
	fmt.Println(x)
	y := isprintable("hi\n")
	fmt.Println(y)
}
