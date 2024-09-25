package main

import "fmt"

// function definition

func isnegative(num int) {
	if num < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}

}

// call and test the funtion
func main() {
	isnegative(6)
	isnegative(55)
	isnegative(-23)
	isnegative(1)
	isnegative(0)
	isnegative(-1)
}
