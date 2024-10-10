package main

import "fmt"

// la concatenation hiya la somme nta3 deux string
// fonction fiha deux parametre de type string et katruturni string
func concat(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	x := concat("my name is", " mehdi")
	fmt.Println(x)
}
