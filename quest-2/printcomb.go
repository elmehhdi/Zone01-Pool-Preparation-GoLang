package main

import "fmt"

func PrintComb() {
	// dernaha triyed f 7 bach tkon i < j < k . par exemple task fiha combinaison nta3 4 nums ghadi ndiroha triyed f 6 .
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				// nprintiw chaque combinaison
				fmt.Printf("%d%d%d", i, j, k)
				// ntchekew wach wselna l 789 si non nprintiw virgule w espace
				if i != 7 || j != 8 || k != 9 {
					fmt.Print(", ") // Add a comma and space except for the last combination
				}
			}
		}
	}
	fmt.Println()
}

// nelghaw l function ntestiwha
func main() {
	PrintComb()
}
