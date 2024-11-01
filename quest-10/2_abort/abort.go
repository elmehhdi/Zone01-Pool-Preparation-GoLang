package main

import "fmt"

func abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	// loop west loop bach n'compariw ila kan i > j nbedlolhom l'blayes
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	// la mediane hiya num li ykon au milieu,
	// le cas fih 5 dyal les numero ghadi ykon howa num li f l'index 2
	return nums[2]
}

func main() {
	median := abort(5, 11, 6, 10, 12)
	fmt.Println(median)

}
