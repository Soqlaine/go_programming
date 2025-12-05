package main

import (
	"fmt"
	"math"
)

func cube(a float64) float64 {
	return math.Pow(a, 3)
}
func main() {

	cube_3 := cube(2)
	fmt.Println(cube_3)
}
