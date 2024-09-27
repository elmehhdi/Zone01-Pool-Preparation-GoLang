package main

import "fmt"

func basicatoi2(s string) int {
	result := 0
	// n'itériw 3la chaque rune
	for _, rune := range s {
		// ila kan rune bin '0' et '9' nconvértéweh si nn nreturniw 0
		if rune >= '0' && rune <= '9' {
			num := int(rune - '0')
			result = result*10 + num

		} else {
			return 0

		}
	}
	return result

}

func main() {
	x := "001023"
	res1 := basicatoi2(x)
	y := "55efh34"
	res2 := basicatoi2(y)
	fmt.Println(res1)
	fmt.Println(res2)

}
