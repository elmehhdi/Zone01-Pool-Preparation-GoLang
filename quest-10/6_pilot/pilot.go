package main

import "fmt"

// struct njem3o fiha les datatypes li f task
type Pilot struct {
	Name     string
	Life     float64 // matekhdemch b int la7e9ach kayna la virgule
	Age      int
	Aircraft int
}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
