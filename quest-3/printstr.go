// - string in Go is essentially a sequence of
// characters (also called runes)
// - a single charactere called a rune in go

package main

import "fmt"

func printstr(s string) {
	// The for-range loop is used for iterating over collections like arrays, slices, maps, and strings.
	// min tebghi t'itéré 3la chi data stucture feker f for-range loop
	for _, char := range s {
		fmt.Printf("%c", char)
	}
	fmt.Println()

}

func main() {
	printstr("hello")
}

// exemple dyal for-range n7esbo fih l vowels f chi string

/*

vowels := 0
s := "hello world"
for _, char := range s {
    if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
        vowels++
    }
}
fmt.Println("Number of vowels:", vowels)

*/

// ---------------------------------------------------------------------------

// exemple kanprintiw fih chaque charactere et l'index nta3o

/*

str := "hello"
for index, char := range str {
    fmt.Printf("Index: %d, Character: %c\n", index, char)
}


*/
