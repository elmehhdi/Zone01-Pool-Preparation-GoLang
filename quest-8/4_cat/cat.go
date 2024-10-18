package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func read_the_file(file_name string) {
	content, err := ioutil.ReadFile(file_name)

	if err != nil {
		fmt.Println("error :", err)
		return
	}
	// a l'aide de la fonction string() nconvertew l contenue dyal lfichier men bytes l string
	fmt.Println(string(content))
}

func main() {
	// nebdaw men l'argument li f l'index 1 7ta l tali , l'index 0 fih smiya dyal l'programme
	args := os.Args[1:]

	for _, file_name := range args {
		read_the_file(file_name)
	}
}
