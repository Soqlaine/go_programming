package main

import "fmt"

type (
	mile      float64
	kilometer float64
)

const m2km = 1.609

func main() {

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Println("The distance between berlin to paris is : ", kmBerlinToParis)

}
