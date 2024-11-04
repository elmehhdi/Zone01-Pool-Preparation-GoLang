package main

import "fmt"

func join(strs []string, sep string) string {
	res := strs[0]

	// n'loopiw 3la les elements li f string w nzido 3lihom separateur
	for i := 1; i < len(strs); i++ {
		res += sep + strs[i]
	}
	return res
}

func main() {
	my_str := []string{"hello", " hi", " hello", " hi"}
	fmt.Println(join(my_str, ":"))
}
