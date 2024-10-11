package main

// ne7tajo os package bach n'accessiw command line argument
import (
	"fmt"
	"os"
)

func main() {
	// os.args howa wa7ed slice dyal les strings w l'element lowel f had slice howa smiya nta3 l'programme
	// nstoriw smiya dyal programme f une variable
	programme_name := os.Args[0]
	// ghadi yprinti ./printprogramme .. pk ? la7e9ach f linux args[0] kaykon fiha le path kamel
	fmt.Println(programme_name)
}
