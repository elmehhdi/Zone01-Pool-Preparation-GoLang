package main

import "fmt"

func addfive(a *int) {
	*a = *a + 5

}

func main() {
	num := 43
	addfive(&num)
	fmt.Println(num)

}
