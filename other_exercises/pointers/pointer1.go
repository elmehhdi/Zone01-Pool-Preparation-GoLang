/*
Without pointers
Explanation :
Original Value in main: We start with number = 5 in main.
Passing a Copy: When we call addTen(number), a copy of number is passed to addTen, so value inside addTen is only a copy of number.
Modification in the Function: Inside addTen, we update value to value + 10, which becomes 15. But since value is just a copy, it doesn’t affect number in main.
Output: The original number in main is still 5, so the output is 5.
*/
package main

import "fmt"

func addten(value int) {
	value = value + 10
}

func main() {
	number := 5

	addten(number)

	fmt.Println("Without pointers:", number)
}
