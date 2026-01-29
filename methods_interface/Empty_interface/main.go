package main

import (
	"fmt"
	"strings"
)

// An empty interface may hold values of any type

// declaring an empty interface
type emptyInterface interface {
}

// declaring a new struct type which has one field of type empty interface
type person struct {
	info interface{}
}

func main() {
	// declaring an empty interface value
	var empty emptyInterface

	// an empty interface may hold values of any type
	// storing an int in the empty interface
	empty = 5
	fmt.Println(empty)

	// storing a string in the empty interface
	empty = "GO"
	fmt.Println(empty)
	fmt.Println(len(empty.(string)))

	//storing a slice in the empty interface
	empty = []int{2, 34, 4}
	fmt.Println(empty)

	// fmt.Println(len(empty)) // error and it does not work

	// retrieving the dynamic value using an assertion
	fmt.Println(len(empty.([]int)))

	fmt.Println(strings.Repeat("*", 10))

	// declaring person value
	you := person{}

	//assigning any value to empty interface field
	you.info = "You name"
	fmt.Println(you.info)

	you.info = 40
	you.info = []float64{4.5, 6., 8.1}

	fmt.Println(you.info)

}
