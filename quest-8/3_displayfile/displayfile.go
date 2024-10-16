package main

import (
	"fmt"
	"io/ioutil" // package bach ne9raw l'content f chi file
	"os"
)

func main() {
	// ila kan 9el men deux arguments neketbo "file name is missing" 3la 7sab task
	if len(os.Args) < 2 {
		fmt.Println("file name is missing")
		return
	}
	// ila kan kter men deux arguments
	if len(os.Args) > 2 {
		fmt.Println("too many arguments")
		return
	}

	// args[0] howa smiya dyal lprogramme , args[1] howa smiya dyal file li bghina l contenue dyalo
	// n'storiweh f une variable
	name_of_file := os.Args[1]

	// had l fonction readfile() kat9ra contenue dyal l'fichier et nstoriw l contenue f la variable content ila makan 7ta error
	content, err := ioutil.ReadFile(name_of_file)
	if err == nil {
		fmt.Println(string(content))
	} else {
		// ila kan chi error nteb3o "error reading file" 3la 7sab task
		fmt.Println("error reading file: ", err)
	}
}
