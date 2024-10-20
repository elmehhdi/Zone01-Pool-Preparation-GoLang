package main

import "fmt"

// n'creyew struct fih deux champs
type point struct {
	x int
	y int
}

// fonction fiha parametre wa7ed de type pointeur kaypointi 3la struct point
func set_point1(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func set_point2(ptr *point) {
	ptr.x = 55
	ptr.y = 66
}

func main() {
	points := &point{}

	set_point1(points)
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)

	set_point2(points)
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)

}
