package main

import "fmt"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		// l point de depart ykon i + 1 bach daymen ykon kbir 3la comb lowla i < j
		for j := i + 1; j <= 99; j++ {
			// "%02d" bach nprintiw 2 num
			fmt.Printf("%02d %02d", i, j)

			// ntchekiw wach rana f lcombinaison talya si nn neprintiw virgule w espace
			if i != 98 || j != 99 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()

}

// nelghaw l fonction
func main() {
	PrintComb2()
}
