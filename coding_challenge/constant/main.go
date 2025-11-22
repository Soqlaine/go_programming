package main

import (
	"fmt"
)

func main() {
	// const daysWeek int= 7
	// const lightSpeed float64= 299792458
	// const pi float64= 3.14159

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	const secPerDay = 60 * 60 * 24
	const daysYear = 365

	fmt.Printf("There are %d seconds in a year.\n", secPerDay*daysYear)

}
