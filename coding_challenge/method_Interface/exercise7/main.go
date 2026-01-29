package main

import "fmt"

type cube struct {
	edge float64
}

// func volume(c float64) float64 {
// 	return c * c * c
// }

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	// x = 4.5
	// v := volume(x.(float64))
	fmt.Printf("Cube Volume: %v\n", v)
}
