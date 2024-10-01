package main

import "fmt"

func recursivepower(num int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	// recursive call ; nelghaw la fonction mera okhra mais b power-1 7ta nweslo 0 bach n'returniw 1
	return num * recursivepower(num, power-1)
}

func main() {
	x := recursivepower(5, 3)
	fmt.Println(x)
}
