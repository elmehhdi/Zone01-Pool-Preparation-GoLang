package main

import (
	"fmt"
	"os"
	"strconv" // package n'convertew bih men ascii l'integer
)

func main() {
	// ila kano les arguments fog 4 manreturniw walo. les arguments : smiya dyal l'programme, num lowel w num zawej m3a l'operateur
	if len(os.Args) != 4 {
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	operator := os.Args[2]
	num2, err2 := strconv.Atoi(os.Args[3])

	// ila kan chi error f num lowel wla zawej manreturniw walo 3la 7ssab task
	if err1 != nil || err2 != nil {
		return
	}

	// n'switshiw 3la 7ssab l'operateur w ila kan l'operateur chi 7aja men ghir + - * / % manreturniw walo 3la 7sab task
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("no division b' 0")
		} else {
			fmt.Println(num1 / num2)
		}
	case "%":
		if num2 == 0 {
			fmt.Println("no modulo b' 0")
		} else {
			fmt.Println(num1 % num2)
		}
	default:
		return
	}
}
