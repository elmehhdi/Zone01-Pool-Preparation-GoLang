package main

import "fmt"

func ultimatedivmod(a *int, b *int) {
	aux := *a / *b // nstoriw la division f wa7ed lvriable temporaire
	*b = *a % *b // nstoriw lmodulo f *b
	*a = aux // dakchi li f aux nredoh l *a

}

// 3lach had lfilm kamel ? ila 7sebna la division directe ghadi nchangew la waleur li f *a

func main() {
	a := 10
	b := 2
	ultimatedivmod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
