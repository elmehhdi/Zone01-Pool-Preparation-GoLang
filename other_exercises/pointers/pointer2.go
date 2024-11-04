/*
With pointers
Explanation
Original Value in main: We start with number = 5 in main.
Passing a Pointer: When we call addTen(&number), we’re passing the address of number, not a copy. Now value in addTen is a pointer to number.
Modification in the Function:
Inside addTen, *value means "go to the original value that value is pointing to and update it."
So, *value = *value + 10 changes the original number in main to 15.
Output: Since we modified the original number through the pointer, number in main is now 15, and the output is 15.
*/
package main

import "fmt"

func addten(value *int) {
	*value = *value + 10
}

func main() {
	number := 5

	addten(&number)

	fmt.Println("With pointers:", number)
}
